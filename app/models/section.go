package models

import (
	"time"
)

type Section struct {
	ID 		     string `gorm:"size:36;not null;unique;primary_key"`
	Name		 string `grom:"size:100"`
	Slug		 string `grom:"size:100"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Categories   []Category
}