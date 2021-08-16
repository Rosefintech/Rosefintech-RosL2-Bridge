/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/5/8 16:20
*/

package contractInterface

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	"math/big"
	zcommon "rosefintech-rosl2/common"
	"rosefintech-rosl2/contractInterface/contract"
	"rosefintech-rosl2/tools"
	"rosefintech-rosl2/tools/sql"
	"strings"
)

//abigen.exe --abi fundManager.abi --pkg contract --type fund --out fundManager.go
//abigen.exe --abi verify.abi --pkg contract --type verify --out verify.go
//abigen.exe --abi schedule.abi --pkg contract --type schedule --out schedule.go

var (
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	//fundManagerFilter *contract.FundFilterer
	//openPositionMutex sync.Mutex
	//outPositionMutex  sync.Mutex
)


func init() {
	var err error
	client, err = ethclient.Dial(zcommon.Conf.Chain.ChainNet)
	if err != nil {
		panic(err)
	}
	key := tools.CancelTargetStr(zcommon.Conf.Chain.PrivateKey, "0x")
	privateKey, err = crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	//tools.Logger.Debug(privateKey.Public())
	chainID = big.NewInt(zcommon.Conf.Chain.ChainID)

	err = sql.FundUserState{}.DeleteAll()
	if err != nil {
		panic(err)
	}
}

func TotalSupply(contractAddr string) (supply *big.Int, err error) {
	IERC20, err := contract.NewERC20(common.HexToAddress(contractAddr), client)
	if err != nil {
		return
	}
	supply, err = IERC20.TotalSupply(nil)
	return
}

func GetErc20Balance(contractAddr, accountAddr string) (balance *big.Int, err error) {
	IERC20, err := contract.NewERC20(common.HexToAddress(contractAddr), client)
	if err != nil {
		return
	}
	balance, err = IERC20.BalanceOf(nil, common.HexToAddress(accountAddr))
	return
}

func GetEthBalance(accountAddr string) (balance *big.Int, err error) {
	balance, err = tools.GetEthBalance(client, accountAddr)
	if err != nil {
		return
	}
	return
}

func WatchOpenPosition(contractAddr string) {
	fundManagerFilter, err := contract.NewFundFilterer(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	ch := make(chan *contract.FundPositionLog)
	_, err = fundManagerFilter.WatchPositionLog(nil, ch)
	if err != nil {
		panic(err)
	}
	for v := range ch {
		sqlData := sql.FundPositionLog{}
		tools.Logger.Debug("watch openPosition:")
		tools.Logger.Debug(v.Addr.String() + " usdc: " + v.UsdcNum.String() + " eth:" + v.EthNum.String())
		account := strings.ToLower(v.Addr.String())
		sqlData.Id = hex.EncodeToString(v.Id[:])
		sqlData.UsdcNum = v.UsdcNum.String()
		sqlData.Hash = v.Raw.TxHash.String()
		sqlData.Timestamp = v.Timestamp.Int64()
		sqlData.Account = account
		sqlData.EthNum = v.EthNum.String()
		err = sqlData.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
			continue
		}
		oldState := sql.FundUserState{
			TotalOpenPosition: "0",
			TotalBrokerage:    "0",
			TotalCharge:       "0",
			TotalOutPosition:  "0",
		}
		err = oldState.Query("account", account)
		if err != nil && err != gorm.ErrRecordNotFound {
			tools.Logger.Error(err.Error())
			continue
		}
		userTotalPosition, ok := new(big.Int).SetString(oldState.TotalOpenPosition, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalOpenPosition)
			continue
		}
		userTotalPosition = userTotalPosition.Add(userTotalPosition, v.UsdcNum)
		oldState.Timestamp = v.Timestamp.Int64()
		oldState.TotalOpenPosition = userTotalPosition.String()

		err = oldState.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
		}
	}
}

func UploadOpenPosition(contractAddr string) {
	fundManagerFilter, err := contract.NewFundFilterer(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	filter, err := fundManagerFilter.FilterPositionLog(nil)
	if err != nil {
		panic(err)
	}
	tools.Logger.Debug("UploadOpenPosition..................................................")
	for filter.Next() {
		sqlData := sql.FundPositionLog{}
		fmt.Println(filter.Event.Addr.String(), ":", filter.Event.UsdcNum.String())
		account := strings.ToLower(filter.Event.Addr.String())
		sqlData.Id = hex.EncodeToString(filter.Event.Id[:])
		sqlData.UsdcNum = filter.Event.UsdcNum.String()
		sqlData.Hash = filter.Event.Raw.TxHash.String()
		sqlData.Timestamp = filter.Event.Timestamp.Int64()
		sqlData.Account = account
		sqlData.EthNum = filter.Event.EthNum.String()
		err = sqlData.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
			continue
		}
		oldState := sql.FundUserState{
			TotalOpenPosition: "0",
			TotalBrokerage:    "0",
			TotalCharge:       "0",
			TotalOutPosition:  "0",
		}
		err = oldState.Query("account", account)
		if err != nil && err != gorm.ErrRecordNotFound {
			tools.Logger.Error(err.Error())
			continue
		}

		userTotalPosition, ok := new(big.Int).SetString(oldState.TotalOpenPosition, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalOpenPosition)
			continue
		}
		userTotalPosition = userTotalPosition.Add(userTotalPosition, filter.Event.UsdcNum)
		oldState.Timestamp = filter.Event.Timestamp.Int64()
		oldState.TotalOpenPosition = userTotalPosition.String()
		err = oldState.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
		}
	}
}


func WatchOutPosition(contractAddr string) {
	fundManagerFilter, err := contract.NewFundFilterer(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	ch := make(chan *contract.FundOutOfPositionLog)
	_, err = fundManagerFilter.WatchOutOfPositionLog(nil, ch)
	if err != nil {
		panic(err)
	}

	for v := range ch {
		sqlData := sql.FundOutOfPositionLog{}
		tools.Logger.Debug("watch WatchOutPosition........")
		tools.Logger.Debug(v.Account.String() + " TokenNum: " + v.TokenNum.String() + " TokenBrokerageNum:" + v.TokenBrokerageNum.String() + " TokenChargeNum:" + v.TokenChargeNum.String())
		account := strings.ToLower(v.Account.String())
		outPositionTransfer(&sqlData, v)
		err = sqlData.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
			continue
		}
		oldState := sql.FundUserState{
			TotalOpenPosition: "0",
			TotalBrokerage:    "0",
			TotalCharge:       "0",
			TotalOutPosition:  "0",
		}
		err = oldState.Query("account", account)
		if err != nil && err != gorm.ErrRecordNotFound {
			tools.Logger.Error(err.Error())
			continue
		}
		totalOutPosition, ok := new(big.Int).SetString(oldState.TotalOutPosition, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalOutPosition)
			continue
		}
		totalCharge, ok := new(big.Int).SetString(oldState.TotalCharge, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalCharge)
			continue
		}
		totalBrokerage, ok := new(big.Int).SetString(oldState.TotalBrokerage, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalBrokerage)
			continue
		}

		totalOutPosition = totalOutPosition.Add(totalOutPosition, v.TokenNum)
		totalCharge = totalCharge.Add(totalCharge, v.TokenChargeNum)
		totalBrokerage = totalBrokerage.Add(totalBrokerage, v.TokenBrokerageNum)

		oldState.Timestamp = v.Timestamp.Int64()
		oldState.TotalOutPosition = totalOutPosition.String()
		oldState.TotalCharge = totalCharge.String()
		oldState.TotalBrokerage = totalBrokerage.String()
		err = oldState.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
		}
	}
}


func UploadOutPosition(contractAddr string) {
	fundManagerFilter, err := contract.NewFundFilterer(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	filter, err := fundManagerFilter.FilterOutOfPositionLog(nil)
	if err != nil {
		panic(err)
	}
	tools.Logger.Debug("UploadOutPosition..................................................")
	for filter.Next() {
		sqlData := sql.FundOutOfPositionLog{}
		fmt.Println(filter.Event.Account, ":", filter.Event.TokenNum, filter.Event.TokenChargeNum, filter.Event.TokenBrokerageNum)
		outPositionTransfer(&sqlData, filter.Event)
		err = sqlData.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
			continue
		}
		oldState := sql.FundUserState{
			TotalOpenPosition: "0",
			TotalBrokerage:    "0",
			TotalCharge:       "0",
			TotalOutPosition:  "0",
		}
		err = oldState.Query("account", sqlData.Account)
		if err != nil && err != gorm.ErrRecordNotFound {
			tools.Logger.Error(err.Error())
			continue
		}

		totalOutPosition, ok := new(big.Int).SetString(oldState.TotalOutPosition, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalOutPosition)
			continue
		}
		totalCharge, ok := new(big.Int).SetString(oldState.TotalCharge, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalCharge)
			continue
		}
		totalBrokerage, ok := new(big.Int).SetString(oldState.TotalBrokerage, 10)
		if !ok {
			tools.Logger.Error("big Int SetString error:" + oldState.TotalBrokerage)
			continue
		}

		//info = zcommon.ChainOutPosition[account]
		totalOutPosition = totalOutPosition.Add(totalOutPosition, filter.Event.TokenNum)
		totalCharge = totalCharge.Add(totalCharge, filter.Event.TokenChargeNum)
		totalBrokerage = totalBrokerage.Add(totalBrokerage, filter.Event.TokenBrokerageNum)
		oldState.Timestamp = filter.Event.Timestamp.Int64()
		oldState.TotalOutPosition = totalOutPosition.String()
		oldState.TotalCharge = totalCharge.String()
		oldState.TotalBrokerage = totalBrokerage.String()
		err = oldState.InsertDB()
		if err != nil {
			tools.Logger.Error(err.Error())
		}
	}
}

func outPositionTransfer(outPositionLog *sql.FundOutOfPositionLog, logOutPositionEvent *contract.FundOutOfPositionLog) {
	outPositionLog.Hash = logOutPositionEvent.Raw.TxHash.String()
	outPositionLog.TokenNum = logOutPositionEvent.TokenNum.String()
	outPositionLog.TokenBrokerageNum = logOutPositionEvent.TokenBrokerageNum.String()
	outPositionLog.TokenChargeNum = logOutPositionEvent.TokenChargeNum.String()
	outPositionLog.Account = strings.ToLower(logOutPositionEvent.Account.String())
	outPositionLog.BillNum = logOutPositionEvent.BillNum.String()
	outPositionLog.TokenAddr = logOutPositionEvent.TokenAddr.String()
	outPositionLog.Timestamp = logOutPositionEvent.Timestamp.Int64()
}


func VerifyWithDraw(contractAddr string, apiUrl string, data *zcommon.WithDraw) (string, bool) {
	var (
		abc           [2][2][2]*big.Int
		input         [21]*big.Int
		User          []common.Address
		ContractAddrs [2]common.Address 
		WithDrawNum   []*big.Int        
		RewardRos     []*big.Int
		BillTokenNum  []*big.Int
		Fee           []*big.Int
		Brokerage     []*big.Int
		LeftEthNumber []*big.Int 
		VrsHash       [4][32]byte
		//PeriodManager common.Address
	)
	if len(data.V) > 64 || len(data.R) > 64 || len(data.S) > 64 || len(data.MsgHash) > 64 {
		return "v r s hash is error", false
	}

	{
		hash := common.Hex2Bytes(data.MsgHash)
		copy(VrsHash[0][32-len(hash):], hash)
		v := common.Hex2Bytes(data.V)
		copy(VrsHash[1][32-len(v):], v)
		r := common.Hex2Bytes(data.R)
		copy(VrsHash[2][32-len(r):], r)
		s := common.Hex2Bytes(data.S)
		copy(VrsHash[3][32-len(s):], s)
	}

	verify, err := contract.NewSchedule(common.HexToAddress(contractAddr), client)
	if err != nil {
		tools.Logger.Error(err.Error())
		return "", false
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				data.ABC[i][j][k] = tools.CancelTargetStr(data.ABC[i][j][k], "0x")
			}
		}
	}

	abc[0][0][0], _ = new(big.Int).SetString(data.ABC[0][0][0], 16)
	abc[0][0][1], _ = new(big.Int).SetString(data.ABC[0][0][1], 16)
	abc[0][1][0], _ = new(big.Int).SetString(data.ABC[1][1][0], 16)
	abc[0][1][1], _ = new(big.Int).SetString(data.ABC[1][1][1], 16)
	abc[1][0][0], _ = new(big.Int).SetString(data.ABC[0][1][0], 16)
	abc[1][0][1], _ = new(big.Int).SetString(data.ABC[0][1][1], 16)
	abc[1][1][0], _ = new(big.Int).SetString(data.ABC[1][0][0], 16)
	abc[1][1][1], _ = new(big.Int).SetString(data.ABC[1][0][1], 16)
	for i, _ := range input {
		data.Input[i] = tools.CancelTargetStr(data.Input[i], "0x")
		input[i], _ = new(big.Int).SetString(data.Input[i], 16)
	}
	for i, _ := range data.User {
		User = append(User, common.HexToAddress(data.User[i]))
	}
	for i, _ := range data.WithDrawNum {
		n, _ := new(big.Int).SetString(data.WithDrawNum[i], 10)
		WithDrawNum = append(WithDrawNum, n)
	}
	for i, _ := range data.RewardRos {
		n, _ := new(big.Int).SetString(data.RewardRos[i], 10)
		RewardRos = append(RewardRos, n)
	}
	for i, _ := range data.BillTokenNum {
		n, _ := new(big.Int).SetString(data.BillTokenNum[i], 10)
		BillTokenNum = append(BillTokenNum, n)
	}
	for i, _ := range data.Fee {
		n, _ := new(big.Int).SetString(data.Fee[i], 10)
		Fee = append(Fee, n)
	}
	for i, _ := range data.Brokerage {
		n, _ := new(big.Int).SetString(data.Brokerage[i], 10)
		Brokerage = append(Brokerage, n)
	}
	for i, _ := range data.LeftEthNumber {
		n, _ := new(big.Int).SetString(data.LeftEthNumber[i], 10)
		LeftEthNumber = append(LeftEthNumber, n)
	}

	ContractAddrs[0] = common.HexToAddress(data.PeriodManager)
	ContractAddrs[1] = common.HexToAddress(data.WithDrawTokenAddr)
	//auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), zcommon.Conf.Chain.PassWord, big.NewInt(3))

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		tools.Logger.Error(err.Error())
		return "", false
	}
	tools.Logger.Debug("Tag=" + data.Tag)
	fmt.Println("PeriodManager=", ContractAddrs[0])
	fmt.Println("User=", User)
	fmt.Println("WithDrawTokenAddr=", ContractAddrs[1])
	fmt.Println("WithDrawNum=", WithDrawNum)
	fmt.Println("RewardRos=", RewardRos)
	fmt.Println("Fee=", Fee)
	fmt.Println("Brokerage=", Brokerage)
	fmt.Println("LeftEthNumber=", LeftEthNumber)
	fmt.Println("abc=", abc)
	fmt.Println("input=", input)
	fmt.Println("v=", VrsHash[0])
	fmt.Println("r=", VrsHash[1])
	fmt.Println("s=", VrsHash[2])
	fmt.Println("hash=", VrsHash[3])
	bytes, err := tools.HttpRequest(apiUrl, "application/json", tools.Get, nil, nil)
	if err != nil {
		tools.Logger.Error(err.Error())
		return err.Error(), false
	}
	obj, err := tools.NewJsonObject(bytes)
	if err != nil {
		tools.Logger.Error(err.Error())
		return err.Error(), false
	}
	gasPrice, err := obj.GetString("result.ProposeGasPrice")
	if err != nil {
		tools.Logger.Error(err.Error())
		return err.Error(), false
	}
	gwei := new(big.Int).SetUint64(1000000000)
	addGas := new(big.Int).SetUint64(5)
	auth.GasPrice, _ = new(big.Int).SetString(gasPrice, 10)
	auth.GasPrice.Add(auth.GasPrice, addGas)
	auth.GasPrice.Mul(auth.GasPrice, gwei)
	chi, err := verify.Withdraw(auth, data.Tag, ContractAddrs, User, WithDrawNum, RewardRos, BillTokenNum, Fee, Brokerage, LeftEthNumber, abc, input, VrsHash)
	if err != nil {
		tools.Logger.Error(err.Error())
		return err.Error(), false
	}
	a := chi.Hash()
	tools.Logger.Debug("chain hash=" + a.String())
	return a.String(), true
}
