package controllers

import (
	"cem-taxi-api/internal/models"
	"cem-taxi-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {
	var mailBody models.MailBody

	if err := c.BindJSON(&mailBody); err != nil {
		log.Printf("error when parsing json body %s", err)
		c.JSON(http.StatusBadRequest, "Body error")
		return
	}

	err := services.SendMail(mailBody)

	if err != "" {
		c.JSON(http.StatusBadRequest, "Mail sending error"+err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
