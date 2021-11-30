package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbAcc *gorm.DB

//Account Kullanıcı hesapları burada olacak.
type Account struct {
	gorm.Model

	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Active    *bool  `json:"active"`
}

func init() {
	config.Connect()
	dbAcc = config.GetDB()
	dbAcc.AutoMigrate(&Account{})
}

//CreateAccount provides
func (acc *Account) CreateAccount() *Account {
	dbAcc.NewRecord(acc)
	dbAcc.Create(&acc)
	return acc
}

//GetAllAccounts provides
func GetAllAccounts() []Account {
	var accounts []Account

	dbAcc.Where("deleted_at is not null").Find(&accounts)
	return accounts
}

//GetAccountByID provides
func GetAccountByID(id int64) (*Account, *gorm.DB) {
	var acc Account

	dbAcc = dbAcc.Where("ID = ?", id).Find(&acc)
	return &acc, dbAcc
}

//DeleteAccount provides
func DeleteAccount(id int64) Account {
	var acc Account

	dbAcc.Where("ID = ?", id).Delete(acc)
	return acc
}
