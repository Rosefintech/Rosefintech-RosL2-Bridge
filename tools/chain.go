/*
@Author :  
@Group :Rosefintech
@Description :
@Data: 2021/5/7 11:45
*/

package tools

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

//gets the current block number
func GetBlockNumber(client *ethclient.Client) (number uint64, err error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	number = header.Number.Uint64()
	return
}

//ETH balance
func GetEthBalance(client *ethclient.Client, account string) (balance *big.Int, err error) {
	header, err := client.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	if err != nil {
		return
	}
	balance = header
	return
}

