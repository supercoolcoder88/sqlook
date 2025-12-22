package orchestrator

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"sqlook/query"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v3"
)

func HandleCommands(ctx context.Context, cmd *cli.Command) error {
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

		rows, err := query.Execute(command, q, db)

		if err != nil {
			return err
		}

		if command == "SELECT" {
			printSelectResponse(rows)
		} else {
			fmt.Printf("%s success", command)
		}
	}

	return nil
}

func printSelectResponse(rows *sql.Rows) error {
	cols, err := rows.Columns()

	if err != nil {
		return err
	}

	vals := make([]any, len(cols))
	scanArgs := make([]any, len(cols))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
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
