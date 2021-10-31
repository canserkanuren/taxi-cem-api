package router

import (
	"cem-taxi-api/internal/controllers"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/go-template/internal/middlewares"
)

func Setup() *gin.Engine {
	app := gin.New()

	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s \"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 - 0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	app.Use(gin.Recovery())

	// Routes
	app.POST("/api/mail", controllers.SendMail)

	app.GET("/api/healthcheck", controllers.Health)

	return app
}
