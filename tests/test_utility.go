package tests

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTestTablesWithData(db *sql.DB) {
	db.Exec(`
		DROP TABLE users;
	`)

	db.Exec(`
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);
	`)

	db.Exec(`
	INSERT INTO users (name, age) VALUES (?, ?)
	`, "John", 30)

	db.Exec(`
	INSERT INTO users (name, age) VALUES (?, ?)
	`, "Jane", 28)

	db.Exec(`
	INSERT INTO users (name, age) VALUES (?, ?)
	`, "Will", 24)

}
