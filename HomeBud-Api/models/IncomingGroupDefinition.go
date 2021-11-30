package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbIngd *gorm.DB

//IncomingGroupDefinition Gelir Grup tanımları burada yapılacak.
type IncomingGroupDefinition struct {
	gorm.Model
	Name   string `json:"name"`
	Active *bool  `json:"active"`
	//Tags   []Tag  `json:"tags"`
}

func init() {
	config.Connect()
	dbIngd = config.GetDB()
	dbIngd.AutoMigrate(&IncomingGroupDefinition{})
}

//CreateIncomingGroupDefinition provides
func (ingd *IncomingGroupDefinition) CreateIncomingGroupDefinition() *IncomingGroupDefinition {
	dbIngd.NewRecord(ingd)
	dbIngd.Create(&ingd)
	return ingd
}

//GetAllIncomingGroupDefinitions provides
func GetAllIncomingGroupDefinitions() []IncomingGroupDefinition {
	var IncomingGroupDefinitions []IncomingGroupDefinition

	dbIngd.Where("deleted_at is not null").Find(&IncomingGroupDefinitions)
	return IncomingGroupDefinitions
}

//GetIncomingGroupDefinitionByID provides
func GetIncomingGroupDefinitionByID(id int64) (*IncomingGroupDefinition, *gorm.DB) {
	var ingd IncomingGroupDefinition

	dbIngd = dbIngd.Where("ID = ?", id).Find(&ingd)
	return &ingd, dbIngd
}

//DeleteIncomingGroupDefinition provides
func DeleteIncomingGroupDefinition(id int64) IncomingGroupDefinition {
	var ingd IncomingGroupDefinition

	dbIngd.Where("ID = ?", id).Delete(ingd)
	return ingd
}
