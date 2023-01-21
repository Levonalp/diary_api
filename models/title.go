package models

import (
	"github.com/jinzhu/gorm"
)

type Title struct {
	gorm.Model
	Name string `json:"name"`
	DiaryID uint `json:"diary_id"`
}

func (t *Title) TableName() string {
	return "titles"
}

func (t *Title) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Title) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *Title) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

func GetTitleByID(db *gorm.DB, id uint) (*Title, error) {
	var title Title
	err := db.First(&title, id).Error
	return &title, err
}

func GetAllTitles(db *gorm.DB) ([]Title, error) {
	var titles []Title
	err := db.Find(&titles).Error
	return titles, err
}

func GetTitlesByDiaryID(db *gorm.DB, diaryID uint) ([]Title, error) {
	var titles []Title
	err := db.Where("diary_id = ?", diaryID).Find(&titles).Error
	return titles, err
}
