package handlers

import (
	"countreez/configs"
	"countreez/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountryIndex(c *gin.Context) {
	countries, err := repositories.GetAllCountries(configs.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "Something went wrong on the server.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Data retieved successfully",
		"data": countries,
	})
}