package models

import (
	config "homebud/config"

	"github.com/jinzhu/gorm"
)

var dbTag *gorm.DB

//Tag Etiketler burada yapılacak.
type Tag struct {
	gorm.Model

	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func init() {
	config.Connect()
	dbTag = config.GetDB()
	dbTag.AutoMigrate(&Tag{})
}

//CreateTag provides
func (t *Tag) CreateTag() *Tag {
	dbTag.NewRecord(t)
	dbTag.Create(&t)
	return t
}

//GetAllTags provides
func GetAllTags() []Tag {
	var tags []Tag

	dbTag.Where("deleted_at is not null").Find(&tags)
	return tags
}

//GetTagByID provides
func GetTagByID(id int64) (*Tag, *gorm.DB) {
	var t Tag

	dbTag = dbTag.Where("ID = ?", id).Find(&t)
	return &t, dbTag
}

//DeleteTag provides
func DeleteTag(id int64) Tag {
	var t Tag

	dbTag.Where("ID = ?", id).Delete(t)
	return t
}
