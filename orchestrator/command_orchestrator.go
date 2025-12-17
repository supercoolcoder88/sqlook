package orchestrator

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v3"
)

func CommandOrchestrator(ctx context.Context, cmd *cli.Command) error {
	// Validate command inputs
	if err := commandInputValidator(cmd); err != nil {
		return err
	}

	fmt.Print("\nAll good\n")

	// Load sqlite db file
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}

	defer db.Close()

	// Query
	q := cmd.String("query")

	if q != "" {
		if err := queryOrchestrator(q, db); err != nil {
			return err
		}
		return nil
	}

	// Error handle

	// Return results

	// TUI
	// Rendering UI
	// List tables

	return nil
}

func queryOrchestrator(query string, db *sql.DB) error {
	// Query orchestration
	// SELECT
	isSelect, err := regexp.MatchString("SELECT", query)

	if err != nil {
		return fmt.Errorf("regex failed: %v", err)
	}

	// TODO : Handle query
	if isSelect {
		// rows, err := db.Query(query)
	}

	// INSERT
	isInsert, err := regexp.MatchString("INSERT", query)

	if err != nil {
		return fmt.Errorf("regex failed: %v", err)
	}

	// TODO: Handle query
	if isInsert {

	}

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("query failed: %v", err)
	}

	return nil
}

func selectQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func commandInputValidator(cmd *cli.Command) error {
	// Must include only file name
	if cmd.NArg() != 1 {
		return fmt.Errorf("ERROR: no file name provided")
	}

	fileName := cmd.Args().Get(0)

	// Must match SQLite file format
	matched, err := regexp.MatchString("[a-z,A-Z,0-9].db", fileName)

	if !matched || err != nil {
		return fmt.Errorf("ERROR: file must end with .db")
	}

	return nil
}
