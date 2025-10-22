package database

import (
	model "cinema-app/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// for sqlserver
	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
	// 	user, pass, host, port, name,
	// )
	// db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	//for postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pass, name, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("db connected!")
	return db, nil
}

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.CinemaBranch{},
		&model.Movie{},
		&model.Showtime{},
		&model.Seat{},
		&model.Transaction{},
		&model.SeatTransaction{},
	)
}
