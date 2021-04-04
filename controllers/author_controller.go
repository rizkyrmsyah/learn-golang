package controllers

import (
	"github.com/rizkyrmsyah/learn-golang/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAuthor(c *gin.Context) {
	var author []models.Author
	err := models.GetAuthor(&author)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, author)
	}
}

func StoreAuthor(c *gin.Context) {
	var author models.Author
	c.BindJSON(&author)
	err := models.AddAuthor(&author)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, author)
	}
}