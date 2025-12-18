package orchestrator

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

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
		if _, err := queryOrchestrator(q, db); err != nil {
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

type QueryResponse struct {
	isInsertSuccess bool
	rows            *sql.Rows
}

func queryOrchestrator(query string, db *sql.DB) (QueryResponse, error) {
	// Query orchestration
	// SELECT
	f := strings.Fields(strings.TrimSpace(query))

	if len(f) == 0 {
		return QueryResponse{}, errors.New("empty query")
	}

	command := strings.ToUpper(f[0])

	// TODO : Handle query
	switch command {
	case "SELECT":
		rows, err := selectQuery(query, db)
		if err != nil {
			return QueryResponse{}, fmt.Errorf("query failed: %v", err)
		}

		return QueryResponse{rows: rows}, nil

	case "INSERT":
		err := insertQuery(query, db)
		if err != nil {
			return QueryResponse{}, fmt.Errorf("query failed: %v", err)
		}

		return QueryResponse{isInsertSuccess: true}, nil

	default:
		fmt.Print("Invalid SQL query")
	}

	return QueryResponse{}, nil
}

func selectQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func insertQuery(query string, db *sql.DB) error {
	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
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
