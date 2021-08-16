/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/5/13 11:36
*/

package contractInterface

//abigen.exe --abi UniswapV2Factory.abi --pkg uniswap --type UniswapV2Factory --out UniswapV2Factory.go
//abigen.exe --abi fundManager.abi --pkg fundManager --type fund --out fundManager.go

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	zcommon "rosefintech-rosl2/common"
	"rosefintech-rosl2/contractInterface/contract"
	"rosefintech-rosl2/tools"
	"strings"
)

var (
	ethNet   *ethclient.Client
	router02 *contract.V2router02
)

//const (
//	uniswapFactory = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
//	uniswapRouter2="0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
//	EthRopsten="wss://ropsten.infura.io/ws/v3/39d9cc578a1647228b9367b0a9253ccc"
//)

func init() {
	var err error
	ethNet, err = ethclient.Dial(zcommon.Conf.Chain.ChainNet)
	if err != nil {
		panic(err.Error())
	}
	router02, err = contract.NewV2router02(common.HexToAddress(zcommon.Conf.Contract.UniswapRouter), ethNet)
	if err != nil {
		panic(err)
	}
}

func Swap(key string, amountIn, amountOutMin int64, pair []string, to string, deadline int64) error {
	var (
		arrPair []common.Address
	)
	if len(pair) < 2 {
		return errors.New("pair address < 2")
	}
	for _, v := range pair {
		arrPair = append(arrPair, common.HexToAddress(v))
	}
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", big.NewInt(3))
	if err != nil {
		return err
	}
	trans, err := router02.SwapExactTokensForTokens(auth, big.NewInt(amountIn), big.NewInt(amountOutMin), arrPair, common.HexToAddress(to), big.NewInt(deadline))
	if err != nil {
		return err
	}
	fmt.Println(trans.Hash())
	return nil
}


func GetAmountsOut(amountIn int64, pair []string) (int64, error) {
	var (
		arrPair []common.Address
	)
	if len(pair) < 2 {
		return 0, errors.New("pair address < 2")
	}
	for _, v := range pair {
		arrPair = append(arrPair, common.HexToAddress(v))
	}
	//arrPair:=

	data, err := router02.GetAmountsOut(nil, big.NewInt(amountIn), arrPair)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("amount out=", data)
	return data[1].Int64(), nil
}


func GetErc20Decimals(erc20Addr common.Address) (number int, err error) {
	erc20, err := contract.NewERC20(erc20Addr, ethNet)
	if err != nil {
		return
	}
	n, err := erc20.Decimals(nil)
	if err != nil {
		return
	}
	number = int(n)
	return
}


func GetReserves(pairAddr common.Address, token0, token1 common.Address) (result struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, err error) {

	uniswapV2Pair, err := contract.NewIUniswapV2Pair(pairAddr, ethNet)
	if err != nil {
		return
	}
	result, err = uniswapV2Pair.GetReserves(nil)
	if err != nil {
		return
	}
	t0, err := uniswapV2Pair.Token0(nil)
	if err != nil {
		return
	}

	if token0 != t0 {
		result.Reserve0, result.Reserve1 = result.Reserve1, result.Reserve0
	}
	return
}

func GetUniswapPair(token1, token2 common.Address) (pairAddr common.Address, err error) {
	uniswapV2Factory, err := contract.NewUniswapV2Factory(common.HexToAddress(zcommon.Conf.Contract.UniswapFactory), ethNet)
	if err != nil {
		return
	}
	pairAddr, err = uniswapV2Factory.GetPair(nil, token1, token2)
	return
}

func GetUniswapPrice(coin, market string) (price float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			tools.Logger.Error(fmt.Sprintf("%v %v", "Recovered in Error:", r))
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("not know panic")
			}
		}
	}()

	pair, err := GetUniswapPair(common.HexToAddress(coin), common.HexToAddress(market))
	if err != nil {
		return 0, err
	}
	//fmt.Println("pair=",pair.String())
	result, err := GetReserves(pair, common.HexToAddress(coin), common.HexToAddress(market))
	if err != nil {
		return 0, err
	}

	token1dec, err := GetErc20Decimals(common.HexToAddress(coin))
	if err != nil {
		return 0, err
	}

	token2dec, err := GetErc20Decimals(common.HexToAddress(market))
	if err != nil {
		return 0, err
	}
	pool1, err := tools.Int64StrToFloat64(result.Reserve0.String(), token1dec)
	if err != nil {
		return 0, err
	}
	pool2, err := tools.Int64StrToFloat64(result.Reserve1.String(), token2dec)
	if err != nil {
		return 0, err
	}
	fmt.Println("pool1=", pool1)
	fmt.Println("pool2=", pool2)

	if pool2 < pool1 {
		price = tools.CutOutFloat64(pool2/pool1, 5) //int(int64(pool1)/int64(pool2)))
	} else {
		price = tools.CutOutFloat64(pool2/pool1, 2)
	}

	return
}
