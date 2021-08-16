/*
@Author :
@Group :Rosefintech
@Description :
@Data: 2021/8/3 18:29
*/

package sql

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./fundLog.db3"
)

var (
	sqlite *gorm.DB
)

func init() {
	db, err := gorm.Open(dbDriverName, dbName)
	if err != nil {
		panic(err)
	}
	sqlite = db
	err = CreateOpenPositionTable()
	if err != nil {
		panic(err)
	}
	err = CreateOutPositionTable()
	if err != nil {
		panic(err)
	}
	err = CreateFundUserStateTable()
	if err != nil {
		panic(err)
	}
}

type FundUserState struct {
	Account           string `gorm:"account"`
	TotalOpenPosition string `gorm:"total_open_position"`
	TotalOutPosition  string `gorm:"total_out_position"`
	TotalCharge       string `gorm:"total_charge"`
	TotalBrokerage    string `gorm:"total_brokerage"` 
	Timestamp         int64  `gorm:"timestamp"`
}

func (FundUserState) TableName() string {
	return "fund_user_state"
}

type FundOutOfPositionLog struct {
	Account           string `gorm:"account"`
	TokenAddr         string `gorm:"token_addr"`
	BillNum           string `gorm:"bill_token_num"`
	TokenNum          string `gorm:"token_num"`
	TokenBrokerageNum string `gorm:"token_brokerage_num"`
	TokenChargeNum    string `gorm:"token_charge_num"`
	Timestamp         int64  `gorm:"timestamp"`
	Hash              string `gorm:"hash"`
}

func CreateOutPositionTable() error {
	sql := `create table if not exists "out_position" (
		"hash" text not null,
		"account" text not null,
		"token_addr" text not null,
		"bill_num" text,
		"token_num" text,
		"token_brokerage_num" text,
		"token_charge_num" text,
		"timestamp" int,
		unique("hash","account")
	)`
	r := sqlite.Exec(sql)
	return r.Error
}

func (FundOutOfPositionLog) TableName() string {
	return "out_position"
}

func (d *FundOutOfPositionLog) InsertDB() error {
	return sqlite.Create(d).Error
}

type FundPositionLog struct {
	Id        string `gorm:"id"`
	Account   string `gorm:"account"`
	EthNum    string `gorm:"eth_num"`
	UsdcNum   string `gorm:"usdc_num"`
	Timestamp int64  `gorm:"timestamp"`
	Hash      string `gorm:"hash"`
}

func CreateOpenPositionTable() error {
	sql := `create table if not exists "open_position" (
		"hash" text not null,
		"id" text unique,
		"account" text not null,
		"eth_num" text,
		"usdc_num" text,
		"timestamp" int
	)`
	r := sqlite.Exec(sql)
	return r.Error
}

func (FundPositionLog) TableName() string {
	return "open_position"
}

func (d *FundPositionLog) InsertDB() error {
	return sqlite.Create(d).Error
}

func CreateFundUserStateTable() error {
	sql := `create table if not exists "fund_user_state" (
		"account" text not null unique,
		"total_open_position" text,
		"total_out_position" text,
		"total_charge" text,
		"total_brokerage" text,
		"timestamp" int
	)`
	r := sqlite.Exec(sql)
	return r.Error
}

func (d *FundUserState) Query(typeName, typeValue string) (err error) {
	if d != nil {
		return errors.New("point is null")
	}
	err = sqlite.Where(typeName+"=?", typeValue).Find(d).Error
	return
}

func (d *FundUserState) InsertDB() (err error) {
	var temp FundUserState
	err = sqlite.Where("account = ?", d.Account).First(&temp).Error
	if err == gorm.ErrRecordNotFound {
		err = sqlite.Create(d).Error
	} else {
		err = sqlite.Model(d).Where("account=?", d.Account).Update(d).Error
	}
	return
}

func (d FundUserState) DeleteAll() error {
	return sqlite.Model(d).Delete(d).Error
}

