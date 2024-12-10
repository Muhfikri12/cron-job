package database

import (
	"cron-job/database/seeder"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	log := &zap.Logger{}
	err := seeder.ExampleSeeder(db)
	if err != nil {
		log.Error("filed to seed example seeder", zap.Error(err))
		return
	}
}
