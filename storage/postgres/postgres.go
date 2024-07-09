package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	//cnf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", "localhost", 5433, "postgres", "payment", "1111")
	return sql.Open("postgres", conDb)
}
