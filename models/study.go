package models

type Study struct {
	ID         uint         `gorm:"primarykey"`
	Name       string       `gorm:"unique;not null"`
	StudyPhoto []StudyPhoto `gorm:"foreignKey:StudyID"`
}

type StudyPhoto struct {
	ID      uint `gorm:"primarykey"`
	StudyID uint
	Photo   string `gorm:"unique;not null"`
}
