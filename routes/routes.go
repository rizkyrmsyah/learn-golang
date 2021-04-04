package routes

import (
	"github.com/rizkyrmsyah/learn-golang/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
 	r := gin.Default()
 	v1 := r.Group("/v1")
	{
		v1.GET("authors", controllers.GetAuthor)
		v1.POST("authors", controllers.StoreAuthor)
		// v1.GET("authors/:id", controllers.GetAuthorByID)
		// v1.PUT("authors/:id", controllers.UpdateAuthor)
		// v1.DELETE("authors/:id", controllers.DeleteAuthor)
	}
 	return r
}