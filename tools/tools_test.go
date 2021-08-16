/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2021/5/7 11:51
*/

package tools

import (
	"fmt"
	"rosefintech-rosl2/tools/sql"
	"testing"
)


func TestInsertFundUserState(t *testing.T) {
	err:=sql.CreateFundUserStateTable()
	fmt.Println(err)
	var data sql.FundUserState
	data.TotalOpenPosition="123456"
	data.Account="0x265a25621a6s51"
	data.TotalBrokerage="556165"
	data.TotalCharge="1113323111"
	err=data.InsertDB()
	fmt.Println(err)
}
