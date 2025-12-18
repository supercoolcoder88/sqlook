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
	res, err := queryOrchestrator("SELECT * FROM users", db)

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

	resInsert, err := queryOrchestrator(`INSERT INTO users (name, age) VALUES ('Mark', 11)`, db)

	// Logic Fix: Error if NOT success OR if err is present
	if err != nil || !resInsert.isInsertSuccess {
		t.Errorf("Insert failed: %v", err)
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

func Test_insertQuery(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")

	defer db.Close()

	tests.CreateEmptyTestTables(db)

	err := insertQuery(`INSERT INTO users (name, age) VALUES ("John", 30)`, db)

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
