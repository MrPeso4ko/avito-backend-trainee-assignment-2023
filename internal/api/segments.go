package api

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"avito-backend-trainee-assignment-2023/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func bindAndValidateSegment(context *gin.Context) (*models.Segment, error) {
	segment := &models.Segment{}
	err := context.BindJSON(segment)
	if err != nil {
		logger.Errorf("Could not bind segment JSON: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "bad json"})
		return nil, err
	}
	//if segment.Name == "" {
	//	logger.Errorf("Got segment without name")
	//	context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "parameter 'segment_name' unfilled"})
	//	return nil, errors.New("parameter 'name' unfilled")
	//}
	logger.Debugf("Segment name: %s", segment.Name)
	return segment, nil
}

func createSegment(db *sqlx.DB, context *gin.Context) {
	logger.Info("Creating segment...")
	segment, err := bindAndValidateSegment(context)
	if err != nil {
		return
	}
	res, err := storage.InsertSegment(db, segment)
	if err != nil {
		logger.Errorf("Could not insert segment: %+v", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal error"})
		return
	} else if res == 0 {
		logger.Info("Segment isn't present in database")
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "segment with this name is already created"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "Segment created successfully"})
	logger.Info("Created")
}

func deleteSegment(db *sqlx.DB, context *gin.Context) {
	logger.Info("Deleting segment...")
	segment, err := bindAndValidateSegment(context)
	if err != nil {
		return
	}
	res, err := storage.DeleteSegment(db, segment)
	if err != nil {
		logger.Errorf("Could not delete segment: %+v", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal error"})
		return
	} else if res == 0 {
		logger.Info("Segment isn't present in database")
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": true, "message": "there is no segment with this name"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Segment deleted"})
	logger.Info("Deleted")
}
