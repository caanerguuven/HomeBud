package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbInd *gorm.DB

//IncomingDefinition Gelir tanımları burada yapılacak.
type IncomingDefinition struct {
	gorm.Model
	Name                      string `json:"name"`
	IncomingDefinitionGroupID *int   `json:"incomeDefinitionGroupId"`
	Active                    *bool  `json:"active"`
}

func init() {
	config.Connect()
	dbInd = config.GetDB()
	dbInd.AutoMigrate(&IncomingDefinition{})
}

//CreateIncomingDefinition provides
func (ind *IncomingDefinition) CreateIncomingDefinition() *IncomingDefinition {
	dbInd.NewRecord(ind)
	dbInd.Create(&ind)
	return ind
}

//GetAllIncomingDefinitions provides
func GetAllIncomingDefinitions() []IncomingDefinition {
	var incomingDefinitions []IncomingDefinition

	dbInd.Where("deleted_at is not null").Find(&incomingDefinitions)
	return incomingDefinitions
}

//GetIncomingDefinitionByID provides
func GetIncomingDefinitionByID(id int64) (*IncomingDefinition, *gorm.DB) {
	var ind IncomingDefinition

	dbInd = dbInd.Where("ID = ?", id).Find(&ind)
	return &ind, dbInd
}

//DeleteIncomingDefinition provides
func DeleteIncomingDefinition(id int64) IncomingDefinition {
	var ind IncomingDefinition

	dbInd.Where("ID = ?", id).Delete(ind)
	return ind
}
