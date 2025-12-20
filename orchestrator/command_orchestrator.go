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

	f := cmd.Args().Get(0)

	// Load sqlite db file
	db, err := sql.Open("sqlite3", f)

	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}

	defer db.Close()

	// Query
	q := cmd.String("query")

	if q != "" {
		f := strings.Fields(strings.TrimSpace(q))

		if len(f) == 0 {
			return errors.New("empty query")
		}

		command := strings.ToUpper(f[0])

		res, err := queryOrchestrator(q, db, command)

		if err != nil {
			fmt.Printf("%v failed\n", command)
			return err
		}

		if command == "SELECT" {
			cols, err := res.rows.Columns()

			if err != nil {
				return err
			}

			vals := make([]any, len(cols))
			scanArgs := make([]any, len(cols))
			for i := range vals {
				scanArgs[i] = &vals[i]
			}

			for res.rows.Next() {
				if err := res.rows.Scan(scanArgs...); err != nil {
					return err
				}

				for i, colName := range cols {
					switch v := vals[i].(type) {
					case nil:
						fmt.Printf("%s: NULL | ", colName)
					case []byte:
						fmt.Printf("%s: %s | ", colName, string(v))
					default:
						fmt.Printf("%s: %v | ", colName, v)
					}
				}
			}
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

type queryResponse struct {
	isSuccess bool
	rows      *sql.Rows
}

func queryOrchestrator(query string, db *sql.DB, command string) (queryResponse, error) {
	// Query orchestration
	switch command {
	case "SELECT":
		rows, err := selectQuery(query, db)
		if err != nil {
			return queryResponse{isSuccess: false}, fmt.Errorf("query failed: %v", err)
		}

		return queryResponse{isSuccess: true, rows: rows}, nil

	case "INSERT", "CREATE":
		err := execQuery(query, db)
		if err != nil {
			return queryResponse{isSuccess: false}, fmt.Errorf("query failed: %v", err)
		}

		return queryResponse{isSuccess: true}, nil

	default:
		return queryResponse{}, fmt.Errorf("error: invalid SQL query")
	}
}

func selectQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func execQuery(query string, db *sql.DB) error {
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
