package models

type Portfolio struct {
	ID             uint             `gorm:"primarykey"`
	Name           string           `gorm:"unique;not null"`
	PortfolioPhoto []PortfolioPhoto `gorm:"foreignKey:PortfolioID"`
}

type PortfolioPhoto struct {
	ID          uint `gorm:"primarykey"`
	PortfolioID uint
	Photo       string `gorm:"unique;not null"`
}
