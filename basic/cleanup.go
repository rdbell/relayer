package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// every hour, delete all very old events
func cleanupRoutine(db *sqlx.DB) {
	for {
		time.Sleep(60 * time.Minute)
		// TODO: query board app for list of abusive users and clean up their old events
		// db.Exec(`DELETE FROM event WHERE created_at < $1`, time.Now().AddDate(0, -3, 0))
	}
}
