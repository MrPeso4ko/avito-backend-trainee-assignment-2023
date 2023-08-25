package storage

import (
	"avito-backend-trainee-assignment-2023/internal/models"
	"github.com/jmoiron/sqlx"
)

func UpsertSegment(db *sqlx.DB, segment models.Segment) error {
	_, err := db.NamedExec(`INSERT INTO segments_manager.segments (name) VALUES (:name) ON CONFLICT (name) DO NOTHING`,
		segment)
	return err
}

func DeleteSegment(db *sqlx.DB, segment models.Segment) error {
	_, err := db.NamedExec(`DELETE FROM segments_manager.segments WHERE name=:name`, segment)
	return err
}
