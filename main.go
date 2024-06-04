package main

import (
	// "database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Data structure for the greeting template
type GreetingData struct {
	Name string
}

// Database initialization function
func initDB() (*sqlx.DB, error) {

	connStr := "host=localhost port=6444 user=xenous password=xenous dbname=xenous sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id uuid PRIMARY KEY,
            name text NOT NULL
        )
    `)
	if err != nil {
		return nil, err
	}
	// ctx := context.Background()
	// db.NamedExecContext(ctx, '', )
	insertUser(db, "Oleg")
	return db, nil
}

// Function to insert a user into the database
func insertUser(db *sqlx.DB, name string) error {
	var err error
	for i := 0; i < 10; i++ {
		id := uuid.New()
		_, err := db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
		fmt.Println(err)
	}
	return err
}

func main() {
	// Initialize the database
	db, err := initDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()
}
