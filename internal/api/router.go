package api

import (
	"avito-backend-trainee-assignment-2023/pkg/logging"
	"avito-backend-trainee-assignment-2023/pkg/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = logging.GetLogger()

func StartAPIServer() {
	logger.Info("Starting API server...")

	db, err := postgres.Connect("postgres://default:0000@172.17.0.1:5432/postgres?sslmode=disable") // TODO: config based on .env
	if err != nil {
		logger.Errorf("error connecting to postgres: %+v", err)
		return
	}

	r := gin.Default()
	err = r.SetTrustedProxies(nil)
	if err != nil {
		logger.Errorf("error setting trusted proxies: %+v", err)
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/segments", func(c *gin.Context) {
		createSegment(db, c)
	})

	r.DELETE("/segments", func(c *gin.Context) {
		deleteSegment(db, c)
	})

	r.GET("/users", func(c *gin.Context) {
		getSegments(db, c)
	})

	r.PUT("/users", func(c *gin.Context) {
		addToSegments(db, c)
	})

	r.DELETE("/users", func(c *gin.Context) {
		removeFromSegments(db, c)
	})

	err = r.Run("0.0.0.0:9053") // TODO: config based on .env
	if err != nil {
		logger.Errorf("Error starting server: %+v", err)
		return
	}
}
