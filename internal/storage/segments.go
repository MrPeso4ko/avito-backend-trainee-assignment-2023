package storage

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

func getSegmentId(db *sqlx.DB, segment *models.Segment) (int, error) {
	var segmentId []int
	err := db.Select(&segmentId, `SELECT id FROM segments_manager.segments WHERE name = $1`, segment.Name)
	if err != nil {
		return 0, err
	}
	if len(segmentId) == 0 {
		return -1, errors.New("no such segment")
	}
	return segmentId[0], err
}

func InsertSegment(db *sqlx.DB, segment *models.Segment) (int64, error) {
	res, err := db.NamedExec(`INSERT INTO segments_manager.segments (name) VALUES (:name) ON CONFLICT (name) DO NOTHING`,
		segment)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteSegment(db *sqlx.DB, segment *models.Segment) (int64, error) {
	_, err := db.NamedExec(`DELETE
FROM segments_manager.belongs_segment
WHERE segment_id = (SELECT id FROM segments_manager.segments WHERE name = :name)`, segment)
	res, err := db.NamedExec(`DELETE FROM segments_manager.segments WHERE name=:name`, segment)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
