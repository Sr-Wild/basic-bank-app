package main

import (
	"time"

	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {

	now := time.Now()
	movements := []struct {
		amount     float64
		balance    float64
		multiplier int
		created_at time.Time
	}{
		{
			amount:     5_000,
			balance:    5_000,
			multiplier: 1,
			created_at: now.AddDate(0, 0, -5),
		},
		{
			amount:     1_500,
			balance:    6_500,
			multiplier: 1,
			created_at: now.AddDate(0, 0, -4),
		},
		{
			amount:     500,
			balance:    6_000,
			multiplier: -1,
			created_at: now.AddDate(0, 0, -3),
		},
		{
			amount:     8_000_000,
			balance:    8_006_000,
			multiplier: 1,
			created_at: now.AddDate(0, 0, -1),
		},
	}

	up := func(db orm.DB) error {
		for _, movement := range movements {
			_, err := db.Exec(`
				INSERT INTO "movement" (user_id, amount, balance, multiplier, created_at) VALUES(?, ?, ?, ?, ?);
			`, userIDuser1, movement.amount, movement.balance, movement.multiplier, movement.created_at)
			if err != nil {
				return err
			}
		}
		return nil
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`DELETE FROM "movement" WHERE user_id = ?`, userIDuser1)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20221207011102_add_movements_user1", up, down, opts)
}
