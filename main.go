package main

import (
	// "database/sql"

	"encoding/json"
	"fmt"

	jopa2 "github.com/KozhurkinTimur/dota2/name"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Data structure for the greeting template
type GreetingData struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
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

	// deleteUser(db)
	// insertUser(db)
	// getUser(db)

	return db, nil
}

// Function to insert a user into the database
func insertUser(db *sqlx.DB) error {
	var err error
	for i := 0; i < 5; i++ {
		name := jopa2.Dota2Name() //вместо этого можно поставить randomdata.name
		id := uuid.New()
		_, err := db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
		fmt.Println(err)
	}
	return err
}

func deleteUser(db *sqlx.DB, inputId uuid.UUID) (error, GreetingData) {
	var err error
	var grData GreetingData
	// fmt.Scanln(&inputId)
	err = db.Get(&grData, "DELETE FROM users WHERE id = $1 RETURNING *", inputId)

	if err != nil {
		fmt.Println(err)
	}

	return err, grData
}

func getUser(db *sqlx.DB) error {
	result, err := db.Queryx("SELECT * FROM users")

	for result.Next() {
		var id uuid.UUID
		var name string

		err = result.Scan(&id, &name)

		fmt.Println("id: " + id.String() + " name: " + name)
	}

	return err
}

func updateUser(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error) {
	var err error
	var grData = new(GreetingData)

	name := "" //jopa2.Dota2Name()
	if name == "" {
		err = db.Get(grData, "SELECT * FROM users WHERE id = $1", inputId)
		fmt.Println(grData.Name)
		if err != nil {
			fmt.Println(err)
		}
		err = db.Get(grData, "UPDATE users SET name = $1 WHERE id = $2 RETURNING *", grData.Name, inputId)
	} else {
		err = db.Get(grData, "UPDATE users SET name = $1 WHERE id = $2 RETURNING *", name, inputId)
	}

	if err != nil {
		fmt.Println(err)
	}
	return grData, err
}

func Jsonify(obj any) string {
	if obj == nil {
		return ""
	}

	str, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}

func main() {
	// Initialize the database
	db, err := initDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	// err, grData := deleteUser(db, uuid.MustParse("d7006448-e348-4d6d-bcf5-4ba4b66af877"))
	grData, err := updateUser(db, uuid.MustParse("2fcb2070-d483-4a98-8b9d-a7d8dd35435c"))

	defer db.Close()

	n := Jsonify(grData)
	fmt.Println(n)
}
