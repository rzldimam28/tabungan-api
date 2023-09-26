package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rzldimam28/tabungan-api/internal/config"
)

func NewDatabaseConnection(cnf *config.Config) *sql.DB {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Name,
	)
	
	connection, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error when connecting database: %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error when ping connection: %s", err.Error())
	}

	return connection

}