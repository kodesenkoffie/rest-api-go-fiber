package blogs

import (
	"github/kodesenkoffie/server/internal/config/db"
	"github/kodesenkoffie/server/pkg/models"
)

func createBlog(blogData models.Blog) (models.Blog, error) {
	if err := db.DB.Create(&blogData).Error; err != nil {
		return models.Blog{}, err
	}
	return blogData, nil
}

func fetchAllBlogs() []models.Blog {
	var blogs []models.Blog
	db.DB.Find(&blogs)
	return blogs
}

func fetchBlogById(id string) models.Blog {
	var blog models.Blog
	db.DB.First(&blog, &id)
	return blog
}

func updateBlog() {}

func deleteBlog() {}
