package storage

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"github.com/jmoiron/sqlx"
)

func GetSegments(db *sqlx.DB, user *models.User) ([]*models.Segment, error) {
	var res []*models.Segment
	err := db.Select(&res, `SELECT name
FROM segments_manager.segments
         JOIN segments_manager.belongs_segment bs ON
    segments.id = bs.segment_id
WHERE bs.user_id = $1`, user.Id)
	if err != nil {
		return nil, err
	}
	return res, err
}

func AddUserToSegment(db *sqlx.DB, user *models.User, segment *models.Segment) (int64, error) {
	segmentId, err := getSegmentId(db, segment)
	if err != nil {
		return int64(segmentId), err
	}
	res, err := db.Exec(`INSERT INTO segments_manager.belongs_segment (user_id, segment_id)
VALUES ($1, $2)
ON CONFLICT (user_id, segment_id) DO NOTHING `, user.Id, segmentId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func RemoveFromSegment(db *sqlx.DB, user *models.User, segment *models.Segment) (int64, error) {
	segmentId, err := getSegmentId(db, segment)
	if err != nil {
		return int64(segmentId), err
	}
	res, err := db.Exec(`DELETE
FROM segments_manager.belongs_segment
WHERE user_id = $1
  AND segment_id = $2`, user.Id, segmentId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
