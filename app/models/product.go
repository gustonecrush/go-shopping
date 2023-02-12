package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID 				 string `gorm:"size:36;not null;unique;primary_key"`
	ParentID		 string `grom:"size:36;index"`
	User			 User
	UserID			 string `gorm:"size:36;index"`
	Sku		         string `gorm:"size:36;index"`
	Name		     string `gorm:"size:255"`
	Slug             string `gorm:"size:255"`
	Price	         decimal.Decimal `gorm:"type:decimal(16,2)"`
	Stock			 int
	Weight		     decimal.Decimal `gorm:"type:decimal(10,2)"`
	ShortDescription string
	Description      string `gorm:"type:text"`
	Status		     int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt		 gorm.DeletedAt
	ProductImages    []ProductImage
}