package services

import (
	"blog/config"
	"blog/models"
)

// GetAllBlogs - Mendapatkan semua blog dari database
func GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := config.DB.Preload("Category").Preload("SubCategory").Preload("Tags").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

// CreateBlog - Membuat blog baru
func CreateBlog(blog models.Blog) (*models.Blog, error) {
	if err := config.DB.Create(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetBlogsByTag - Mendapatkan blog berdasarkan tag
func GetBlogsByTag(tagName string) ([]models.Blog, error) {
	var blogs []models.Blog
	if err := config.DB.Preload("Category").
		Preload("SubCategory").
		Preload("Tags").
		Joins("JOIN blog_tags ON blog_tags.blog_id = blogs.id").
		Joins("JOIN tags ON tags.id = blog_tags.tag_id").
		Where("tags.name = ?", tagName).
		Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetBlogByID - Mendapatkan blog berdasarkan ID
func GetBlogByID(id uint) (*models.Blog, error) {
	var blog models.Blog
	if err := config.DB.Preload("Category").Preload("SubCategory").Preload("Tags").First(&blog, id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// UpdateBlog - Memperbarui blog berdasarkan ID
func UpdateBlog(id uint, updatedBlog models.Blog) (*models.Blog, error) {
	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	// Perbarui data blog
	blog.Title = updatedBlog.Title
	blog.Content = updatedBlog.Content
	blog.CategoryID = updatedBlog.CategoryID
	blog.SubCategoryID = updatedBlog.SubCategoryID
	blog.Tags = updatedBlog.Tags

	if err := config.DB.Save(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// DeleteBlog - Menghapus blog berdasarkan ID
func DeleteBlog(id uint) error {
	if err := config.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}
