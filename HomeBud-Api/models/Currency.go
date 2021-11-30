package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbCur *gorm.DB

//Currency para birimleri burada gösterilecek.
type Currency struct {
	gorm.Model

	Name   string `json:"name"`
	Active *bool  `json:"active"`
}

func init() {
	config.Connect()
	dbCur = config.GetDB()
	dbCur.AutoMigrate(&Currency{})
}

//CreateCurrency provides
func (cur *Currency) CreateCurrency() *Currency {
	dbCur.NewRecord(cur)
	dbCur.Create(&cur)
	return cur
}

//GetAllCurrencies provides
func GetAllCurrencies() []Currency {
	var currencies []Currency

	dbCur.Where("deleted_at is not null").Find(&currencies)
	return currencies
}

//GetCurrencyByID provides
func GetCurrencyByID(id int64) (*Currency, *gorm.DB) {
	var cur Currency

	dbCur = dbCur.Where("ID = ?", id).Find(&cur)
	return &cur, dbCur
}

//DeleteCurrency provides
func DeleteCurrency(id int64) Currency {
	var cur Currency

	dbCur.Where("ID = ?", id).Delete(cur)
	return cur
}
