package models

import (
	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	URL string `json:"url"`
	DiaryID uint `json:"diary_id"`
}

func (p *Photo) TableName() string {
	return "photos"
}

func (p *Photo) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Photo) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Photo) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func GetPhotoByID(db *gorm.DB, id uint) (*Photo, error) {
	var photo Photo
	err := db.First(&photo, id).Error
	return &photo, err
}

func GetAllPhotos(db *gorm.DB) ([]Photo, error) {
	var photos []Photo
	err := db.Find(&photos).Error
	return photos, err
}

func GetPhotosByDiaryID(db *gorm.DB, diaryID uint) ([]Photo, error) {
	var photos []Photo
	err := db.Where("diary_id = ?", diaryID).Find(&photos).Error
	return photos, err
}
