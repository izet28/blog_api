package controllers

import (
	"blog/models"
	"blog/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBlogs - Handler untuk mengambil semua blog
func GetBlogs(c *gin.Context) {
	blogs, err := services.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// CreateBlog - Handler untuk membuat blog baru
func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBlog, err := services.CreateBlog(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBlog)
}

// GetBlogsByTag - Handler untuk mengambil blog berdasarkan tag
func GetBlogsByTag(c *gin.Context) {
	tagName := c.Param("tag")
	blogs, err := services.GetBlogsByTag(tagName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(blogs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No blogs found for this tag"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// UpdateBlog - Handler untuk memperbarui blog berdasarkan ID
func UpdateBlog(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBlog models.Blog
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog, err := services.UpdateBlog(uint(id), updatedBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// DeleteBlog - Handler untuk menghapus blog berdasarkan ID
func DeleteBlog(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.DeleteBlog(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
