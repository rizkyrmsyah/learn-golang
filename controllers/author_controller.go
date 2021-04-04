package controllers

import (
	"github.com/rizkyrmsyah/learn-golang/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func getAuthor(c *gin.Context) {
	var author []models.Author
	err := models.GetAuthor(&author)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, author)
	}
}

				