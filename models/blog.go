package models

import (
	"gorm.io/gorm"
)

// Model Blog
type Blog struct {
	gorm.Model
	Title         string      `json:"title" gorm:"type:varchar(100);not null"`
	Content       string      `json:"content" gorm:"type:text;not null"`
	CategoryID    uint        `json:"category_id" gorm:"not null" `                     // Foreign Key untuk Category
	Category      Category    `json:"category" gorm:"constraint:OnUpdate:CASCADE;"`     // Tidak ada OnDelete
	SubCategoryID uint        `json:"sub_category_id" gorm:"not null"`                  // Foreign Key untuk SubCategory
	SubCategory   SubCategory `json:"sub_category" gorm:"constraint:OnUpdate:CASCADE;"` // Tidak ada OnDelete
	Tags          []Tag       `json:"tags" gorm:"many2many:blog_tags;"`
}
