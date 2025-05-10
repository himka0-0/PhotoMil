package models

type Rental struct {
	ID          uint          `gorm:"primarykey"`
	Name        string        `gorm:"unique;not null"`
	RentalPhoto []RentalPhoto `gorm:"foreignKey:RentalID"`
}

type RentalPhoto struct {
	ID       uint `gorm:"primarykey"`
	RentalID uint
	Photo    string `gorm:"unique;not null"`
}
