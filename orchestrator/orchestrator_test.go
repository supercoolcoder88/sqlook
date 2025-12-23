package orchestrator

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"os"
	"sqlook/query"
	"sqlook/tests"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v3"
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

func Test_HandleCommands_SelectSuccess(t *testing.T) {
	ctx := context.Background()

	tmpDB, err := os.CreateTemp("", "test_*.db")
	if err != nil {
		t.Fatal(err)
	}
	dbPath := tmpDB.Name()
	defer os.Remove(dbPath)
	tmpDB.Close()

	db, _ := sql.Open("sqlite3", dbPath)
	_, err = db.Exec("CREATE TABLE users (id INTEGER, name TEXT); INSERT INTO users VALUES (1, 'Alice');")
	if err != nil {
		t.Fatalf("failed setup: %v", err)
	}
	db.Close()

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "query"},
		},
	}

	args := []string{"executable", dbPath}

	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String("query", "SELECT * FROM users", "")

	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		return HandleCommands(ctx, c)
	}

	err = cmd.Run(ctx, append(args, "--query", "SELECT * FROM users"))

	if err != nil {
		t.Errorf("HandleCommands failed: %v", err)
	}
}
