⚠️ Important: This library uses CGO. You must have a C compiler (like gcc) installed and CGO_ENABLED=1 in your environment for it to work.
1. Installation
Bash

go get github.com/mattn/go-sqlite3

2. Basic Skeleton

Since this follows the standard database/sql interface, you don't use the library directly after the import. You interact with the sql package.
Go

package main

import (
	"database/sql"
	"log"

	// Register the sqlite3 driver.
	// We use the blank identifier (_) because we don't use the package directly.
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open creates a database file 'data.db' if it doesn't exist
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

3. Creating Tables & Inserting Data

Use db.Exec for commands that don't return rows (CREATE, INSERT, UPDATE, DELETE).
Go

func initDB(db *sql.DB) {
	// 1. Create a table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// 2. Insert data safely using placeholders (?)
	// Always use ? to prevent SQL Injection
	insertStmt := `INSERT INTO users (name, age) VALUES (?, ?)`
	
	_, err := db.Exec(insertStmt, "Alice", 30)
	if err != nil {
		log.Printf("Insert error: %v", err)
	}
    
    // You can get the LastInsertId if needed
    res, _ := db.Exec(insertStmt, "Bob", 25)
    id, _ := res.LastInsertId()
    log.Printf("Added Bob with ID: %d", id)
}

4. Querying Data

Use db.Query for multiple rows and db.QueryRow for single rows.
Go

func getUsers(db *sql.DB) {
	// 1. Run the query
	rows, err := db.Query("SELECT id, name, age FROM users WHERE age > ?", 20)
	if err != nil {
		log.Fatal(err)
	}
	// Always close rows to release the connection back to the pool
	defer rows.Close()

	// 2. Iterate over the results
	for rows.Next() {
		var id int
		var name string
		var age int

		// Map columns to variables
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		log.Printf("User: %d | %s | %d", id, name, age)
	}

	// 3. Check for iteration errors
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

Common Gotchas

    Concurrency: SQLite accepts one writer at a time. The driver handles this, but if you have high concurrency, you might see "database is locked" errors. You can mitigate this by enabling Write-Ahead Logging (WAL) mode:
    Go

db.Exec("PRAGMA journal_mode=WAL;")