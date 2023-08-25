package api

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"avito-backend-trainee-assignment-2023/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func createSegment(db *sqlx.DB, context *gin.Context) {
	segment := models.Segment{}
	logger.Info("Creating segment...")
	err := context.BindJSON(&segment)
	if err != nil {
		logger.Errorf("Could not bind segment JSON: %+v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Segment isn't created"})
	}
	logger.Debugf("Segment name: %s", segment.Name)

	err = storage.UpsertSegment(db, segment)
	if err != nil {
		logger.Errorf("Could not bind segment JSON: %+v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Segment isn't created"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Segment created successfully"})
}
