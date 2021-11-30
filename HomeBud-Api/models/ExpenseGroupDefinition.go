package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbExgd *gorm.DB

//ExpenseGroupDefinition Gider grup tanımları burada yapılacak.
type ExpenseGroupDefinition struct {
	gorm.Model

	Name   string `json:"name"`
	Active *bool  `json:"active"`
	//Tags   []Tag  `json:"tags"`
}

func init() {
	config.Connect()
	dbExgd = config.GetDB()
	dbExgd.AutoMigrate(&ExpenseGroupDefinition{})
}

//CreateExpenseGroupDefinition provides
func (exgd *ExpenseGroupDefinition) CreateExpenseGroupDefinition() *ExpenseGroupDefinition {
	dbExgd.NewRecord(exgd)
	dbExgd.Create(&exgd)
	return exgd
}

//GetAllExpenseGroupDefinitions provides
func GetAllExpenseGroupDefinitions() []ExpenseGroupDefinition {
	var ExpenseGroupDefinitions []ExpenseGroupDefinition

	dbExgd.Where("deleted_at is not null").Find(&ExpenseGroupDefinitions)
	return ExpenseGroupDefinitions
}

//GetExpenseGroupDefinitionByID provides
func GetExpenseGroupDefinitionByID(id int64) (*ExpenseGroupDefinition, *gorm.DB) {
	var exgd ExpenseGroupDefinition

	dbExgd = dbExgd.Where("ID = ?", id).Find(&exgd)
	return &exgd, dbExgd
}

//DeleteExpenseGroupDefinition provides
func DeleteExpenseGroupDefinition(id int64) ExpenseGroupDefinition {
	var exgd ExpenseGroupDefinition

	dbExgd.Where("ID = ?", id).Delete(exgd)
	return exgd
}
