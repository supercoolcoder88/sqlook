package orchestrator

import (
	"bytes"
	"database/sql"
	"sqlook/query"
	"sqlook/tests"
	"testing"
)

func Test_PrintSelectResponseSuccess(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	tests.CreateEmptyTestTables(db)

	_, err = query.Execute(
		"INSERT",
		`INSERT INTO users (name, age) VALUES ("John", 50)`,
		db,
	)
	if err != nil {
		t.Fatalf("INSERT failed: %v", err)
	}

	rows, err := query.Execute(
		"SELECT",
		`SELECT name, age FROM users`,
		db,
	)
	if err != nil {
		t.Fatalf("SELECT failed: %v", err)
	}
	defer rows.Close()

	var buf bytes.Buffer

	err = printSelectResponse(&buf, rows)
	if err != nil {
		t.Fatalf("printSelectResponse failed: %v", err)
	}

	got := buf.String()
	want := `| name: John | | age: 50 | `

	if got != want {
		t.Errorf("unexpected output:\nwant: %q\ngot:  %q", want, got)
	}

}
