package services

import (
	"blog/config"
	"blog/models"
	"fmt"
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
	if err := config.DB.Preload("Tags").First(&blog, id).Error; err != nil {
		return nil, err
	}

	// Hapus relasi lama di tabel penghubung (blog_tags)
	if err := config.DB.Model(&blog).Association("Tags").Clear(); err != nil {
		return nil, fmt.Errorf("failed to clear old tags: %v", err)
	}

	// Memeriksa dan membuat atau menemukan tag yang sesuai
	for i, tag := range updatedBlog.Tags {
		if tag.ID == 0 {
			if err := config.DB.FirstOrCreate(&updatedBlog.Tags[i], models.Tag{Name: tag.Name}).Error; err != nil {
				return nil, fmt.Errorf("failed to create or find tag: %s", tag.Name)
			}
		}
	}

	// Memperbarui data blog kecuali relasi tags
	blog.Title = updatedBlog.Title
	blog.Content = updatedBlog.Content
	blog.CategoryID = updatedBlog.CategoryID
	blog.SubCategoryID = updatedBlog.SubCategoryID

	// Simpan perubahan pada blog (kecuali tags)
	if err := config.DB.Save(&blog).Error; err != nil {
		return nil, err
	}

	// Tambahkan relasi tags baru
	if err := config.DB.Model(&blog).Association("Tags").Replace(updatedBlog.Tags); err != nil {
		return nil, fmt.Errorf("failed to update tags: %v", err)
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
