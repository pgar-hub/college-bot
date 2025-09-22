package database

import (
	"database/sql"
	"fmt"
)

func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("sql open:%w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping:%w", err)
	}

	return db, err
}
func CreateTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS schedule")
	if err != nil {
		return fmt.Errorf("drop table: %w", err)
	}
	schedule := `
	CREATE TABLE IF NOT EXISTS schedule (
	id SERIAL PRIMARY KEY,
	week_type VARCHAR(100) NOT NULL,
	day_of_week VARCHAR(100) NOT NULL,
	subject  TEXT NOT NULL,
	cabinet VARCHAR(100) 
	);
	`
	_, err = db.Exec(schedule)
	if err != nil {
		return fmt.Errorf("create schedule table:%w", err)
	}
	fmt.Println("Shedule table created")
	return nil
}
