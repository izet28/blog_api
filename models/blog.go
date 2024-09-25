package models

import (
	"gorm.io/gorm"
)

// Model Blog
type Blog struct {
	gorm.Model
	Title         string      `json:"title" gorm:"type:varchar(100);not null"`
	Content       string      `json:"content" gorm:"type:text;not null"`
	CategoryID    uint        `json:"category_id" gorm:"not null"` // Foreign Key ke Category
	Category      Category    `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubCategoryID uint        `json:"sub_category_id" gorm:"not null" ` // Foreign Key ke SubCategory
	SubCategory   SubCategory `json:"sub_category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tags          []Tag       `json:"tags" gorm:"many2many:blog_tags;"`
}

// Model Category
type Category struct {
	ID            uint          `gorm:"primaryKey"`
	Name          string        `json:"name" gorm:"type:varchar(50);not null;unique"`
	SubCategories []SubCategory `json:"subcategories" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Model SubCategory
type SubCategory struct {
	ID         uint     `gorm:"primaryKey"`
	Name       string   `json:"name" gorm:"type:varchar(50);not null;unique"`
	CategoryID uint     `json:"category_id" gorm:"not null"` // Foreign Key untuk Category
	Category   Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
