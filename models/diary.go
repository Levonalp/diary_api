package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Diary struct {
	gorm.Model
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Photos  []Photo   `gorm:"foreignkey:DiaryID"`
	Date    time.Time `json:"date"`
}

func (d *Diary) TableName() string {
	return "diaries"
}

func (d *Diary) Create(db *gorm.DB) error {
	return db.Create(d).Error
}

func (d *Diary) Update(db *gorm.DB) error {
	return db.Save(d).Error
}

func (d *Diary) Delete(db *gorm.DB) error {
	return db.Delete(d).Error
}

func GetDiaryByID(db *gorm.DB, id uint) (*Diary, error) {
	var diary Diary
	err := db.First(&diary, id).Error
	return &diary, err
}

func GetAllDiaries(db *gorm.DB) ([]Diary, error) {
	var diaries []Diary
	err := db.Find(&diaries).Error
	return diaries, err
}
