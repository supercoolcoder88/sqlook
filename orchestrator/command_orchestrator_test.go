package orchestrator

import (
	"database/sql"
	"sqlook/tests"
	"testing"
)

func Test_queryOrchestrator(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	tests.CreateTestTablesWithData(db)

	// Select query should return contain values
	res, err := queryOrchestrator("SELECT * FROM users", db, "SELECT")

	if err != nil {
		t.Errorf("error: %v", err)
	}

	count := 0
	for res.rows.Next() {
		count += 1
	}

	if count != 3 {
		t.Errorf("expected 3 got %d", count)
	}

	defer res.rows.Close()

	resInsert, err := queryOrchestrator(`INSERT INTO users (name, age) VALUES ('Mark', 11)`, db, "INSERT")

	if err != nil || !resInsert.isSuccess {
		t.Errorf("insert failed: %v", err)
	}
}

func Test_selectQuery(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	tests.CreateTestTablesWithData(db)

	rows, err := selectQuery("SELECT * FROM users", db)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	count := 0
	for rows.Next() {
		count += 1
	}

	if count != 3 {
		t.Errorf("expected 3 got %d", count)
	}

	defer rows.Close()
}

func Test_insert(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	tests.CreateEmptyTestTables(db)

	err := execQuery(`INSERT INTO users (name, age) VALUES ("John", 30)`, db)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	rows, err := db.Query(`SELECT * FROM users`)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	count := 0
	for rows.Next() {
		count += 1
	}

	if count != 1 {
		t.Errorf("expected 1 got %d", count)
	}

	defer rows.Close()
}

func Test_createTable(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	err := execQuery(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);`, db)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	err = execQuery(`INSERT INTO users (name, age) VALUES ("John", 30)`, db)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	rows, err := db.Query(`SELECT * FROM users`)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	count := 0
	for rows.Next() {
		count += 1
	}

	if count != 1 {
		t.Errorf("expected 1 got %d", count)
	}

	defer rows.Close()
}
