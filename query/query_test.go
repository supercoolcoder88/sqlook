package query

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_Execute(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	_, err := Execute("CREATE", `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER
		)
	`, db)
	if err != nil {
		t.Fatalf("CREATE failed: %v", err)
	}

	_, err = Execute("INSERT", `INSERT INTO users (name, age) VALUES ("John", 30)`, db)
	if err != nil {
		t.Fatalf("INSERT failed: %v", err)
	}

	rows, err := Execute("SELECT", `SELECT name, age FROM users`, db)
	if err != nil {
		t.Fatalf("SELECT failed: %v", err)
	}
	if rows == nil {
		t.Fatalf("expected rows, got nil")
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			t.Fatalf("scan failed: %v", err)
		}
		count++
	}
	if count != 1 {
		t.Fatalf("expected 1 row, got %d", count)
	}

	_, err = Execute("UPDATE", `UPDATE users SET name = "Updated", age = 99 WHERE id = 1`, db)
	if err != nil {
		t.Fatalf("UPDATE failed: %v", err)
	}

	rows, err = Execute("SELECT", `SELECT name, age FROM users WHERE id = 1`, db)
	if err != nil {
		t.Fatalf("verification SELECT failed: %v", err)
	}
	if rows == nil {
		t.Fatalf("expected rows, got nil")
	}
	defer rows.Close()

	var name string
	var age int
	if rows.Next() {
		if err := rows.Scan(&name, &age); err != nil {
			t.Fatalf("scan failed: %v", err)
		}
	}

	if name != "Updated" {
		t.Errorf("expected name 'Updated', got %q", name)
	}
	if age != 99 {
		t.Errorf("expected age 99, got %d", age)
	}
}
