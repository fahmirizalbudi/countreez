package handlers

import (
	"countreez/configs"
	"countreez/repositories"
	"countreez/structs"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CountryIndex(c *gin.Context) {
	cacheKey := "countries:all"

	cached, err := configs.Redis.Get(configs.Ctx, cacheKey).Result()
	if err == nil {
		var countries []structs.Country
		if err := json.Unmarshal([]byte(cached), &countries); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Data retrieved successfully (from redis)",
				"data":    countries,
			})
			return
		}
	}

	countries, err := repositories.GetAllCountries(configs.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "Something went wrong on the server.",
		})
		return
	}

	json, _ := json.Marshal(countries)
	configs.Redis.Set(configs.Ctx, cacheKey, json, 1*time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Data retieved successfully",
		"data": countries,
	})
}