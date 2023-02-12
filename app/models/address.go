package models

import "time"

type Address struct {
	ID 		   string `gorm:"size:36;not null;unique;primary_key`
	User       User
	UserID     string `gorm:"size:36;index`
	Name       string `gorm:"size:100"`
	IsPrimary  bool
	Address1   string `gorm:"size:100"`
	Address2   string `gorm:"size:100"`
	Phone      string `gorm:"size:100"`
	Email      string `gorm:"size:100"`
	CityID     string `gorm:"size:255"`
	ProvinceID string `gorm:"size:255"`
	PostCode   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}