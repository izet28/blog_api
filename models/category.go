package models

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
