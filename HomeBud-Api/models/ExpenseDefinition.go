package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbExd *gorm.DB

//ExpenseDefinition Gider tanımları burada yapılacak.
type ExpenseDefinition struct {
	gorm.Model

	Name                     string `json:"name"`
	ExpenseDefinitionGroupID *int   `json:"expenseDefinitionGroupId"`
	Active                   *bool  `json:"active"`
}

func init() {
	config.Connect()
	dbExd = config.GetDB()
	dbExd.AutoMigrate(&ExpenseDefinition{})
}

//CreateExpenseDefinition provides
func (exd *ExpenseDefinition) CreateExpenseDefinition() *ExpenseDefinition {
	dbExd.NewRecord(exd)
	dbExd.Create(&exd)
	return exd
}

//GetAllExpenseDefinitions provides
func GetAllExpenseDefinitions() []ExpenseDefinition {
	var expenseDefinitions []ExpenseDefinition

	dbExd.Where("deleted_at is not null").Find(&expenseDefinitions)
	return expenseDefinitions
}

//GetExpenseDefinitionByID provides
func GetExpenseDefinitionByID(id int64) (*ExpenseDefinition, *gorm.DB) {
	var exd ExpenseDefinition

	dbExd = dbExd.Where("ID = ?", id).Find(&exd)
	return &exd, dbExd
}

//DeleteExpenseDefinition provides
func DeleteExpenseDefinition(id int64) ExpenseDefinition {
	var exd ExpenseDefinition

	dbExd.Where("ID = ?", id).Delete(exd)
	return exd
}
