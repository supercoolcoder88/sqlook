package orchestrator

import (
	"database/sql"
	"sqlook/tests"
	"testing"
)

func Test_queryOrchestrator(t *testing.T) {

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
