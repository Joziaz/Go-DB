package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	user     = "user"
	password = "josias1228"
	hostname = "localhost"
	port     = "5432"
)

var (
	once sync.Once
	db   *sql.DB
)

func dataSourceName(database string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		hostname, port, user, password, database,
	)
}

func GetConnection() *sql.DB {
	var err error
	once.Do(func() {
		db, err = sql.Open("pgx", dataSourceName("DB"))
		if err != nil {
			log.Fatalf("Error %s, when open the database", err)
		}
		pingErr := db.Ping()

		if pingErr != nil {
			log.Fatalf("Error %s, when open the database", err)
		}
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Second)
	})
	return db

}
