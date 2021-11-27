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
		c.IndentedJSON(http.StatusBadRequest, "Body error")
		return
	}

	err := services.SendMail(mailBody)

	if err != "" {
		c.IndentedJSON(http.StatusBadRequest, "Mail sending error"+err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
	return
}
