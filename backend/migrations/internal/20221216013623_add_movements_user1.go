package main

import (
	"math/rand"
	"time"

	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		balance := 8_006_000
		multiplier := 1
		minAmount := 5_000
		maxAmount := 900_000
		for i := 0; i < 1000; i++ {
			amount := rand.Intn(maxAmount-minAmount) + minAmount
			if i%3 == 0 {
				multiplier = -1
			} else {
				multiplier = 1
			}

			balance += amount * multiplier
			_, err := db.Exec(`
				INSERT INTO "movement" (user_id, amount, balance, multiplier, created_at) VALUES(?, ?, ?, ?, ?);
			`, userIDuser1, amount, balance, multiplier, time.Now())

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

	migrations.Register("20221216013623_add_movements_user1", up, down, opts)
}
