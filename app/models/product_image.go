package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductImage struct {
	ID 				 string `gorm:"size:36;not null;unique;primary_key"`
	Product			 Product
	ProductID	     string `gorm:"size:36;index"`
	Path		     string `gorm:"type:text"`
	ExtraLarge		 string `gorm:"type:text"`
	Large            string `gorm:"type:text"`
	Medium	         string `gorm:"type:text"`
	Small			 int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}