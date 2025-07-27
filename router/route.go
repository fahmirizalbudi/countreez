package router

import (
	"countreez/handlers"

	"github.com/gin-gonic/gin"
)

func CountreezRoute() *gin.Engine  {
	r := gin.Default()

	r.GET("api/countries", handlers.CountryIndex)

	return r
}