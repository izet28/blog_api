package controllers

import (
	"blog/models"
	"blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory - Handler untuk membuat category baru
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory, err := services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

// CreateSubCategory - Handler untuk membuat subcategory baru
func CreateSubCategory(c *gin.Context) {
	var subCategory models.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSubCategory, err := services.CreateSubCategory(subCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSubCategory)
}
