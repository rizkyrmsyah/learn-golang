package routes

import (
	"github.com/rizkyrmsyah/learn-golang/controllers/author_controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
 	r := gin.Default()
 	authors := r.Group("/authors")
	{
		authors.GET("authors", author_controller.GetAuthor)
		// authors.POST("authors", controllers.CreateAuthor)
		// authors.GET("authors/:id", controllers.GetAuthorByID)
		// authors.PUT("authors/:id", controllers.UpdateAuthor)
		// authors.DELETE("authors/:id", controllers.DeleteAuthor)
	}
 	return r
}