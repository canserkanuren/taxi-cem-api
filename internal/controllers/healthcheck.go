package controllers

import (
	"cem-taxi-api/pkg/health"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-template/pkg/health"
)

func Health(c *gin.Context) {
	check := health.Check{
		AppName: "Cem Taxi API",
		Version: "1.0.0",
		Status:  "OK",
	}

	c.JSON(http.StatusOK, check)
}
