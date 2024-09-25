package models

// Model Tag
type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
}
