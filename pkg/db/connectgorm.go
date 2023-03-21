package db

import (

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/fazilnbr/project-workey/pkg/config"
	domain "github.com/fazilnbr/project-workey/pkg/domain"
)

func ConnectGormDB(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	psqlInfo := cfg.DBSOURCE
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&domain.User{},
		&domain.Profile{},
		&domain.Verification{},
		&domain.Category{},
		&domain.Job{},
		&domain.Favorite{},
		&domain.Address{},
		&domain.Request{},
		&domain.JobPayment{},
	)

	return db, dbErr
}
