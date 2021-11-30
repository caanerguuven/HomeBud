package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbMoer *gorm.DB

//MonthlyExpenseRecord Aylık Giderler burada yapılacak.
type MonthlyExpenseRecord struct {
	gorm.Model

	Year                *int     `json:"year"`
	Month               *int     `json:"month"`
	ExpenseDefinitionID *int     `json:"expenseDefinitionId"`
	EstimatedAmount     *float64 `json:"estimatedAmount"`
	ActualAmount        *float64 `json:"actualAmount"`
	CurrencyID          *int     `json:"currencyId"`
	Description         string   `json:"description"`
	IsPaid              bool     `json:"isPaid"`
	CreatedAccountID    *int     `json:"createdAccountId"`
}

func init() {
	config.Connect()
	dbMoer = config.GetDB()
	dbMoer.AutoMigrate(&MonthlyExpenseRecord{})
}

//CreateMonthlyExpenseRecord provides
func (moer *MonthlyExpenseRecord) CreateMonthlyExpenseRecord() *MonthlyExpenseRecord {
	moer.IsPaid = false
	dbMoer.NewRecord(moer)
	dbMoer.Create(&moer)
	return moer
}

//GetAllMonthlyExpenseRecords provides
func GetAllMonthlyExpenseRecords() []MonthlyExpenseRecord {
	var monthlyExpenseRecords []MonthlyExpenseRecord

	dbMoer.Where("deleted_at is not null").Find(&monthlyExpenseRecords)
	return monthlyExpenseRecords
}

//GetMonthlyExpenseRecordByID provides
func GetMonthlyExpenseRecordByID(id int64) (*MonthlyExpenseRecord, *gorm.DB) {
	var moer MonthlyExpenseRecord

	dbMoer = dbMoer.Where("ID = ?", id).Find(&moer)
	return &moer, dbMoer
}

//DeleteMonthlyExpenseRecord provides
func DeleteMonthlyExpenseRecord(id int64) MonthlyExpenseRecord {
	var moer MonthlyExpenseRecord

	dbMoer.Where("ID = ?", id).Delete(moer)
	return moer
}
