package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// every hour, delete all very old events
func cleanupRoutine(db *sqlx.DB) {
	for {
		// TODO: query board app for list of abusive users and clean up their old events?
		time.Sleep(60 * time.Minute)
	}
}
