/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/4/28 9:37
*/

package common

import (
	"math/big"
	"sync"
)


type VerifyData struct {
	UserAddress               string `json:"userAddress"`
	Tag                       string `json:"tag"`
	UsdcNumber                string `json:"usdcNumber"`
	Fee                       string `json:"fee"`
	Brokerage                 string `json:"brokerage"`
	RewardRos                 string `json:"rewardRos"`
	DestroySwift              string `json:"destroySwift"`
	LeftEthNumber             string `json:"leftEthNumber"`
	LeftAssets                string `json:"leftAssets"`
	PrevCalcAssets            string `json:"prevCalcAssets"`
	WbtcPrice                 string `json:"wbtcPrice"`
	EthPrice                  string `json:"ethPrice"`
	FundManageContractAddress string `json:"fundManageContractAddress"`
	TestBuildCost             string `json:"buildCost"`
	V                         string `json:"v"`
	R                         string `json:"r"`
	S                         string `json:"s"`
	MsgHash                   string `json:"msgHash"`
}

type WithDraw struct {
	Tag               string
	PeriodManager     string
	User              []string
	WithDrawTokenAddr string
	WithDrawNum       []string
	RewardRos         []string
	BillTokenNum      []string
	Fee               []string
	Brokerage         []string
	LeftEthNumber     []string
	ABC               [2][2][2]string
	Input             [21]string
	V                 string
	R                 string
	S                 string
	MsgHash           string
}

type ZkVerifyData struct {
	TotalBalance           [3]string
	TotalBill              string
	WithdrawDetail         [20]string
	WithdrawBrokerage      [20]string
	WithdrawCharge         [20]string
	UserBillBalance        [20]string
	UserWithdrawBillNum    [20]string
	UserTotalBalance       [20]string
	UserAlreadyWithdraw    [20]string
	UserHistoryBrokerage   [20]string
	UserTotalInvestHistory [20]string
	Price                  [20][2]string

}

type OutPosition struct {
	TokenNum          *big.Int `gorm:"token_num"`
	TokenChargeNum    *big.Int `gorm:"token_charge_num"`
	TokenBrokerageNum *big.Int `gorm:"token_brokerage_num"`
}

type OpenPosition struct {
	TotalOpenPosition big.Int
}

type ProofInfo struct {
	A []string   `json:"a"`
	B [][]string `json:"b"`
	C []string   `json:"c"`
}

type Proof struct {
	ProofData ProofInfo `json:"proof"`
	Inputs    []string  `json:"inputs"`
}

var (
	Conf *Config
	Lock sync.Mutex
)

type contract struct {
	FundManager    []string `yaml:"fundManager"`
	TargetCoin     []string `yaml:"targetCoin"`
	Schedule       string   `yaml:"schedule"`
	UniswapRouter  string   `yaml:"uniswapRouter"`
	UniswapFactory string   `yaml:"uniswapFactory"`
	Wbtc           string   `yaml:"wbtc"`
	Weth           string   `yaml:"weth"`
	USDC           string   `yaml:"usdc"`
	BillToken      string   `yaml:"billToken"`
}

type chain struct {
	ChainNet    string `yaml:"chainNet"`
	Keystore    string `yaml:"keystore"`
	PassWord    string `yaml:"password"`
	ApiGasPrice string `yaml:"apiGasPrice"`
	ChainID     int64  `yaml:"chainID"`
	PrivateKey  string `yaml:"privateKey"`
}

// Config is kato config.
type Config struct {
	SeverPort string   `yaml:"SeverPort"`
	Debug     bool     `yaml:"Debug"`
	Chain     chain    `yaml:"Chain"`
	Contract  contract `yaml:"Contract"`
}
