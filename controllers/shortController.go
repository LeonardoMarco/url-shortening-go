package controllers

import (
	"labs/shortener/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestInputUrl struct {
	Url string `json:"url"`
}

type RequestUri struct {
	Hash string `uri:"hash"`
}

func PostUrl(c *gin.Context) {
	var newUrl RequestInputUrl

	if err := c.BindJSON(&newUrl); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "an error occurred in creation a shortener"})
	}

	urlExist, error := models.VerifyUrl(newUrl.Url)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "an error occurred in creation a shortener"})
	}

	if urlExist != "" {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": urlExist})
		return
	}

	urlCreated, error := models.InsertUrl(newUrl.Url)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "an error occurred in creation a shortener"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": urlCreated})
}

func RedirectUrl(c *gin.Context) {
	var hash RequestUri
	if err := c.ShouldBindUri(&hash); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	urlRedirect, error := models.GetUrlRedirect(hash.Hash)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Url not found"})
	}

	c.Redirect(http.StatusFound, urlRedirect)
}
