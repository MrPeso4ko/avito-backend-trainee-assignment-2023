package api

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"avito-backend-trainee-assignment-2023/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func bindUser(context *gin.Context) (*models.User, error) {
	user := &models.User{}
	err := context.BindQuery(user)
	if err != nil {
		logger.Errorf("Could not bind user URI: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "bad params"})
		return nil, err
	}
	logger.Debugf("User id: %d", user.Id)
	return user, nil
}

type segmentsResponse struct {
	Err      bool              `json:"error"`
	Segments []*models.Segment `json:"segments"`
}

func getSegments(db *sqlx.DB, context *gin.Context) {
	logger.Info("getting segments...")
	user, err := bindUser(context)
	if err != nil {
		logger.Errorf("Could not get segments for user: %+v", err)
		return
	}
	segments, err := storage.GetSegments(db, user)
	if err != nil {
		logger.Errorf("Can't get user segments: %+v", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal error"})
		return
	}
	context.JSON(http.StatusOK, segmentsResponse{false, segments})
	logger.Info("success")
}

type segmentsBody struct {
	Segments []*models.Segment `json:"segments"`
}

func bindAndValidateSegmentsBody(context *gin.Context) (*segmentsBody, error) {
	segments := &segmentsBody{}
	err := context.BindJSON(segments)
	if err != nil {
		logger.Errorf("Could not bind segments JSON: %+v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "bad json"})
		return nil, err
	}
	logger.Debugf("Segments: %+v", segments)
	return segments, nil
}
func addToSegments(db *sqlx.DB, context *gin.Context) {
	user, err := bindUser(context)
	if err != nil {
		logger.Errorf("Could not get segments for user: %+v", err)
		return
	}
	segments, err := bindAndValidateSegmentsBody(context)
	if err != nil {
		logger.Errorf("Could not get segments for user: %+v", err)
	}
	for _, segment := range segments.Segments {
		res, err := storage.AddUserToSegment(db, user, segment)
		if err != nil {
			logger.Errorf("Couldn't add user %d to segment %s: %+v", user.Id, segment.Name, err)
			if res == -1 {
				context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": true, "message": fmt.Sprintf("segment %s not found", segment.Name)})
			} else {
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal error"})
			}
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Added user to segments"})
}
