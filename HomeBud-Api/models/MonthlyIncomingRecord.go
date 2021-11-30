package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbMoir *gorm.DB

//MonthlyIncomingRecord Aylık Gelen Ödemeleri gösterecek.
type MonthlyIncomingRecord struct {
	gorm.Model

	Year                 *int     `json:"year"`
	Month                *int     `json:"month"`
	IncomingDefinitionID *int     `json:"incomingDefinitionId"`
	EstimatedAmount      *float64 `json:"estimatedAmount"`
	ActualAmount         *float64 `json:"actualAmount"`
	CurrencyID           *int     `json:"currencyId"`
	Description          string   `json:"description"`
	IsTaken              bool     `json:"isTaken"`
}

func init() {
	config.Connect()
	dbMoir = config.GetDB()
	dbMoir.AutoMigrate(&MonthlyIncomingRecord{})
}

//CreateMonthlyIncomingRecord provides
func (moir *MonthlyIncomingRecord) CreateMonthlyIncomingRecord() *MonthlyIncomingRecord {
	moir.IsTaken = false

	dbMoir.NewRecord(moir)
	dbMoir.Create(&moir)
	return moir
}

//GetAllMonthlyIncomingRecords provides
func GetAllMonthlyIncomingRecords() []MonthlyIncomingRecord {
	var monthlyIncomingRecords []MonthlyIncomingRecord

	dbMoir.Where("deleted_at is not null").Find(&monthlyIncomingRecords)
	return monthlyIncomingRecords
}

//GetMonthlyIncomingRecordByID provides
func GetMonthlyIncomingRecordByID(id int64) (*MonthlyIncomingRecord, *gorm.DB) {
	var moir MonthlyIncomingRecord

	dbMoir = dbMoir.Where("ID = ?", id).Find(&moir)
	return &moir, dbMoir
}

//DeleteMonthlyIncomingRecord provides
func DeleteMonthlyIncomingRecord(id int64) MonthlyIncomingRecord {
	var moir MonthlyIncomingRecord

	dbMoir.Where("ID = ?", id).Delete(moir)
	return moir
}
