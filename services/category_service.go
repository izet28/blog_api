package services

import (
	"blog/config"
	"blog/models"
)

// CreateCategory - Membuat category baru
func CreateCategory(category models.Category) (*models.Category, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateSubCategory - Membuat subcategory baru
func CreateSubCategory(subCategory models.SubCategory) (*models.SubCategory, error) {
	// Memeriksa apakah category_id valid
	var category models.Category
	if err := config.DB.First(&category, subCategory.CategoryID).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Create(&subCategory).Error; err != nil {
		return nil, err
	}
	return &subCategory, nil
}
