/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/4/28 9:37
*/

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	zcommon "rosefintech-rosl2/common"
	"rosefintech-rosl2/contractInterface"
	"rosefintech-rosl2/tools"
	"rosefintech-rosl2/tools/sql"
	"strconv"
	"strings"
)

func getZkTotalTargetBalance(zkData *[3]string, targetCoinAddr []string, fundManager []string) error {
	var (
		total [3]*big.Int
	)
	if len(zkData) != len(fundManager) {
		return errors.New("array is not inequality")
	}
	for i := 0; i < 3; i++ {
		total[i] = new(big.Int)
	}
	for i := 0; i < 3; i++ {

		for j := 0; j < 2; j++ {
			balance, err := contractInterface.GetErc20Balance(targetCoinAddr[j], fundManager[i])
			if err != nil {
				return err
			}
			if j == 1 {
				total[j+1] = total[j+1].Add(total[j+1], balance)
			} else {
				total[j] = total[j].Add(total[j], balance)
			}

		}
		balance, err := contractInterface.GetEthBalance(fundManager[i])
		if err != nil {
			return err
		}
		total[1] = total[1].Add(total[1], balance)
	}
	for i, v := range total {
		zkData[i] = v.String()
	}
	return nil
}

func getBillTokenBalance(verifyData []zcommon.VerifyData, zkData *[20]string) error {
	for i, v := range verifyData {
		balance, err := contractInterface.GetErc20Balance(zcommon.Conf.Contract.BillToken, v.UserAddress)
		if err != nil {
			return err
		}
		if 0 == balance.Cmp(big.NewInt(0)) {
			a, _ := new(big.Int).SetString("18446744073709551615", 10)
			zkData[i] = a.String()
			continue
		}
		zkData[i] = balance.String()
	}
	return nil
}

//verify and generate the required zero-knowledge parameters
func verifyAndGenerateZkParameters(verifyData []zcommon.VerifyData) (zkData *zcommon.ZkVerifyData, err error) {
	zkData = new(zcommon.ZkVerifyData)
	//
	totalBillNum, err := contractInterface.TotalSupply(zcommon.Conf.Contract.BillToken)
	if err != nil {
		return
	}
	zkData.TotalBill = totalBillNum.String()
	//
	err = getZkTotalTargetBalance(&zkData.TotalBalance, zcommon.Conf.Contract.TargetCoin, zcommon.Conf.Contract.FundManager)
	if err != nil {
		return
	}
	//
	err = getBillTokenBalance(verifyData, &zkData.UserBillBalance)
	if err != nil {
		return
	}
	for i, v := range verifyData {
		zkData.WithdrawCharge[i] = v.Fee
		//
		zkData.WithdrawDetail[i] = v.UsdcNumber
		//
		zkData.WithdrawBrokerage[i] = v.Brokerage
		//
		zkData.UserWithdrawBillNum[i] = v.DestroySwift
		//
		zkData.UserTotalBalance[i] = v.PrevCalcAssets
		//
		userAddr := strings.ToLower(v.UserAddress)
		state := sql.FundUserState{}
		err = state.Query("account", userAddr)
		if err != nil {
			num := new(big.Int)
			tools.Logger.Debug("user addr=" + userAddr + " TokenNum=" + state.TotalOutPosition)
			totalOutPosition, ok := new(big.Int).SetString(state.TotalOutPosition, 10)
			if !ok {
				tools.Logger.Error("big Int SetString error:" + state.TotalOutPosition)
				continue
			}
			totalCharge, ok := new(big.Int).SetString(state.TotalCharge, 10)
			if !ok {
				tools.Logger.Error("big Int SetString error:" + state.TotalCharge)
				continue
			}
			totalBrokerage, ok := new(big.Int).SetString(state.TotalBrokerage, 10)
			if !ok {
				tools.Logger.Error("big Int SetString error:" + state.TotalBrokerage)
				continue
			}
			num = num.Add(totalOutPosition, totalCharge)
			num = num.Add(totalBrokerage, num)
			zkData.UserAlreadyWithdraw[i] = num.String()
			zkData.UserHistoryBrokerage[i] = totalBrokerage.String()
			tools.Logger.Debug("user addr=" + userAddr + " total=" + num.String())
			tools.Logger.Debug("user addr=" + userAddr + " UserHistoryBrokerage=" + zkData.UserHistoryBrokerage[i])
			tools.Logger.Debug("user addr=" + userAddr + " TokenNum=" + totalOutPosition.String() + " charge=" + totalCharge.String() + " Brokerage=" + totalBrokerage.String())
		} else {
			zkData.UserAlreadyWithdraw[i] = "0"
			zkData.UserHistoryBrokerage[i] = "0"
			tools.Logger.Debug("user addr=" + userAddr + " UserAlreadyWithdraw = 0")
			tools.Logger.Debug("user addr=" + userAddr + " UserHistoryBrokerage = 0")
		}
		if err != nil {
			tools.Logger.Debug(userAddr + " ChainOpenPosition " + ":" + state.TotalOpenPosition)
			zkData.UserTotalInvestHistory[i] = state.TotalOpenPosition
		} else {
			zkData.UserTotalInvestHistory[i] = "0"
			tools.Logger.Debug("UserTotalInvestHistory = 0")
		}
		zkData.Price[i][0] = v.WbtcPrice
		zkData.Price[i][1] = v.EthPrice
	}
	return
}


func arrayTiled(array [20][2]string) []string {
	var strs []string
	for i := 0; i < 20; i++ {
		for j := 0; j < 2; j++ {
			strs = append(strs, array[i][j])
		}
	}
	return strs
}


func spliceString(data *zcommon.ZkVerifyData) string {
	var tstr string
	eth, err := tools.ReserveDecimals(data.TotalBalance[1], 1000000000000)
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	data.TotalBalance[1] = eth
	str1 := strings.Join(data.TotalBalance[:], " ")
	bill, err := tools.ReserveDecimals(data.TotalBill, 1000000000000)
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	str2 := bill
	str3 := strings.Join(data.WithdrawDetail[:], " ")
	str4 := strings.Join(data.WithdrawBrokerage[:], " ")
	str5 := strings.Join(data.WithdrawCharge[:], " ")
	for i, _ := range data.UserBillBalance {
		if data.UserBillBalance[i] == "18446744073709551615" {
			continue
		}
		billBalance, err := tools.ReserveDecimals(data.UserBillBalance[i], 1000000000000)
		if err != nil {
			tools.Logger.Error(err.Error())
		}
		data.UserBillBalance[i] = billBalance
	}
	str6 := strings.Join(data.UserBillBalance[:], " ")
	for i, _ := range data.UserWithdrawBillNum {
		billWithdraw, err := tools.ReserveDecimals(data.UserWithdrawBillNum[i], 1000000000000)
		if err != nil {
			tools.Logger.Error(err.Error())
		}
		data.UserWithdrawBillNum[i] = billWithdraw
	}
	str7 := strings.Join(data.UserWithdrawBillNum[:], " ")
	str8 := strings.Join(data.UserTotalBalance[:], " ")
	str9 := strings.Join(data.UserAlreadyWithdraw[:], " ")
	str10 := strings.Join(data.UserHistoryBrokerage[:], " ")
	str11 := strings.Join(data.UserTotalInvestHistory[:], " ")
	price := arrayTiled(data.Price)
	str12 := strings.Join(price, " ")
	tstr = strings.Join([]string{str1, str2, str3, str4, str5, str6, str7, str8, str9, str10, str11, str12}, " ")
	return tstr
}


func FindMaxAndMin(data []zcommon.VerifyData) (maxBtc, minBtc, maxEth, minEth *big.Int) {
	max, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	maxBtc = new(big.Int)
	maxEth = new(big.Int)
	minBtc = max
	minEth = max
	for _, v := range data {
		btcPrice, ok := new(big.Int).SetString(v.WbtcPrice, 10)
		if !ok {
			btcPrice = new(big.Int)
			tools.Logger.Error("new big int SetString error:" + v.WbtcPrice)
		}
		if 1 == btcPrice.Cmp(maxBtc) {
			maxBtc = btcPrice
		}
		if 1 == minBtc.Cmp(btcPrice) {
			minBtc = btcPrice
		}
		ethPrice, ok := new(big.Int).SetString(v.EthPrice, 10)
		if !ok {
			ethPrice = new(big.Int)
			tools.Logger.Error("new big int SetString error:" + v.EthPrice)
		}
		if 1 == ethPrice.Cmp(maxEth) {
			maxEth = ethPrice
		}
		if 1 == minEth.Cmp(ethPrice) {
			minEth = ethPrice
		}
	}
	return
}


func verifyPriceRange(nowBtcPrice, nowEthPrice, userMaxBtcPrice, userMinBtcPrice, userMaxEthPrice, userMinEthPrice, btcScope, ethScope *big.Int) bool {
	if nowBtcPrice.Add(nowBtcPrice, btcScope).Cmp(userMaxBtcPrice) < 0 && nowBtcPrice.Sub(nowBtcPrice, btcScope).Cmp(userMinBtcPrice) > 0 {
		return false
	}
	if nowEthPrice.Add(nowEthPrice, ethScope).Cmp(userMaxEthPrice) < 0 && nowEthPrice.Sub(nowEthPrice, ethScope).Cmp(userMinEthPrice) > 0 {
		return false
	}
	return true
}

func getCoinPrice(coin, market string) (price *big.Int, err error) {
	priceFloat, err := contractInterface.GetUniswapPrice(coin, market)
	if err != nil {
		return
	}
	tools.Logger.Debug("price=" + strconv.FormatFloat(priceFloat, 'f', -1, 64))
	priceInt := int64(priceFloat * 1000000)
	price = new(big.Int).SetInt64(priceInt)
	return
}

//authentication service
func Verify(c *gin.Context) {
	var (
		originVerifyData []zcommon.VerifyData
		verData          []zcommon.VerifyData
		zkChainData      zcommon.WithDraw
	)
	zcommon.Lock.Lock()
	defer zcommon.Lock.Unlock()
	bytes, err := c.GetRawData()
	if err != nil {
		SendJSON(http.StatusOK, false, err, c)
		return
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &verData)
	if err != nil {
		SendJSON(http.StatusOK, false, err, c)
		return
	}
	originVerifyData = verData
	if len(verData) > 20 || len(verData) == 0 {
		SendJSON(http.StatusOK, false, nil, c)
		return
	}
	if len(verData) != 20 {
		for i := len(verData); i < 20; i++ {
			verData = append(verData, zcommon.VerifyData{
				UserAddress: "0x0000000000000000000000000000000000000000",
				Tag:         "0", UsdcNumber: "0", Fee: "0", Brokerage: "0", RewardRos: "0",
				DestroySwift: "0", LeftAssets: "0", WbtcPrice: "0", EthPrice: "0", LeftEthNumber: "0",
				PrevCalcAssets: "0", TestBuildCost: "0"})
		}
	}
	////check the price range
	//maxBtc, MinBtc, maxEth, minEth := FindMaxAndMin(verData)
	////
	//btcPrice, err := getCoinPrice(zcommon.Conf.Contract.Wbtc,zcommon.Conf.Contract.USDC)
	//if err != nil {
	//	SendJSON(http.StatusOK, false, "Price query failed", c)
	//	return
	//}
	//ethPrice, err := getCoinPrice(zcommon.Conf.Contract.Weth,zcommon.Conf.Contract.USDC)
	//if err != nil {
	//	SendJSON(http.StatusOK, false, "Price query failed", c)
	//	return
	//}
	//btcScope := new(big.Int).SetInt64(1000000000)
	//ethScope := new(big.Int).SetInt64(300000000)
	//if btcPrice != nil && ethPrice != nil && maxBtc != nil && MinBtc != nil && maxEth != nil && minEth != nil {
	//	ok := verifyPriceRange(btcPrice, ethPrice, maxBtc, MinBtc, maxEth, minEth, btcScope, ethScope)
	//	if !ok {
	//		SendJSON(http.StatusOK, false, "price is not range", c)
	//		return
	//	}
	//}
	//

	verifyFun := func([]zcommon.VerifyData) ([]byte, error) {
		data, err := verifyAndGenerateZkParameters(verData)
		if err != nil {
			return nil, err
		}
		parStr := spliceString(data)
		tools.Logger.Debug("parameter=" + parStr)
		_, err = tools.ComputeWitness("abi.json", parStr, "out", "witness")
		if err != nil {
			return nil, err
		}
		proofBytes, err := tools.GenerateProof("out", "g16")
		return proofBytes, err
	}
	proofBytes, err := verifyFun(verData)
	if err != nil {
		err = sql.FundUserState{}.DeleteAll()
		if err != nil {
			tools.Logger.Error(err.Error())
		}
		for _, v := range zcommon.Conf.Contract.FundManager {
			contractInterface.UploadOpenPosition(v)
			contractInterface.UploadOutPosition(v)
		}
		proofBytes, err = verifyFun(verData)
		if err != nil {
			tools.Logger.Error(err.Error())
			SendJSON(http.StatusOK, false, err.Error(), c)
			return
		}
	}

	var proofData zcommon.Proof

	err = json.Unmarshal(proofBytes, &proofData)
	if err != nil {
		tools.Logger.Error(err.Error())
		SendJSON(http.StatusOK, false, err.Error(), c)
		return
	}

	for _, v := range originVerifyData {

		zkChainData.User = append(zkChainData.User, v.UserAddress)
		zkChainData.LeftEthNumber = append(zkChainData.LeftEthNumber, v.LeftEthNumber)
		zkChainData.Brokerage = append(zkChainData.Brokerage, v.Brokerage)
		zkChainData.Fee = append(zkChainData.Fee, v.Fee)
		zkChainData.BillTokenNum = append(zkChainData.BillTokenNum, v.DestroySwift)
		zkChainData.WithDrawNum = append(zkChainData.WithDrawNum, v.UsdcNumber)
		zkChainData.RewardRos = append(zkChainData.RewardRos, v.RewardRos)
	}
	zkChainData.PeriodManager = verData[0].FundManageContractAddress
	zkChainData.Tag = verData[0].Tag
	zkChainData.WithDrawTokenAddr = zcommon.Conf.Contract.USDC
	zkChainData.ABC[0][0][0] = proofData.ProofData.A[0]
	zkChainData.ABC[0][0][1] = proofData.ProofData.A[1]
	zkChainData.ABC[0][1][0] = proofData.ProofData.B[0][0]
	zkChainData.ABC[0][1][1] = proofData.ProofData.B[0][1]
	zkChainData.ABC[1][0][0] = proofData.ProofData.B[1][0]
	zkChainData.ABC[1][0][1] = proofData.ProofData.B[1][1]
	zkChainData.ABC[1][1][0] = proofData.ProofData.C[0]
	zkChainData.ABC[1][1][1] = proofData.ProofData.C[1]
	zkChainData.R = tools.CancelTargetStr(originVerifyData[0].R, "0x")
	zkChainData.S = tools.CancelTargetStr(originVerifyData[0].S, "0x")
	zkChainData.V = tools.CancelTargetStr(originVerifyData[0].V, "0x")
	zkChainData.MsgHash = tools.CancelTargetStr(originVerifyData[0].MsgHash, "0x")
	for i, v := range proofData.Inputs {
		zkChainData.Input[i] = v
	}
	fmt.Println("leftEth=", zkChainData.LeftEthNumber)
	fmt.Println(zkChainData)

	hash, ok := contractInterface.VerifyWithDraw(zcommon.Conf.Contract.Schedule, zcommon.Conf.Chain.ApiGasPrice, &zkChainData)
	tools.Logger.Debug("contract verify is " + strconv.FormatBool(ok))

	SendJSON(http.StatusOK, ok, hash, c)
	return
}


func checkWithdrawValue(values []string) error {
	max, _ := new(big.Int).SetString("", 10)

	for _, v := range values {
		num, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return errors.New(v + " is not bigInt")
		}
		contrast := num.Cmp(max)
		if contrast > 0 {
			return errors.New("the user withdrew too much money")
		}
	}
	return nil
}
