package postgres

import (
	"avito-backend-trainee-assignment-2023/pkg/logging"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var logger = logging.GetLogger()

func Connect(dsn string) (*sqlx.DB, error) {
	logger.Info("Trying to connect to postgres...")
	logger.Debugf("dsn: %s", dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening storage: %s", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %s", err)
	}

	db.SetConnMaxIdleTime(0)
	db.SetConnMaxLifetime(0)

	logger.Info("connected to postgres")
	return db, nil
}
