package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Group route untuk blog
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.GET("/", controllers.GetBlogs)
		blogRoutes.POST("/", controllers.CreateBlog)
		blogRoutes.GET("/tag/:tag", controllers.GetBlogsByTag)
		blogRoutes.PUT("/:id", controllers.UpdateBlog)    // Rute untuk memperbarui blog berdasarkan ID
		blogRoutes.DELETE("/:id", controllers.DeleteBlog) // Rute untuk menghapus blog berdasarkan ID
	}

	return router
}
