package orchestrator

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
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
			fmt.Fprintf(cmd.Writer, "query failed")
			return err
		}

		if command == "SELECT" {
			printSelectResponse(cmd.Writer, rows)
			return nil
		}

		fmt.Fprintf(cmd.Writer, "%s success", command)
	}

	return nil
}

func printSelectResponse(w io.Writer, rows *sql.Rows) error {
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
				fmt.Fprintf(w, "| %s: NULL | ", colName)
			case []byte:
				fmt.Fprintf(w, "| %s: %s | ", colName, string(v))
			default:
				fmt.Fprintf(w, "| %s: %v | ", colName, v)
			}
		}
	}
	return nil
}

func commandInputValidator(cmd *cli.Command) error {
	if cmd.NArg() != 1 {
		return errors.New("usage: <command> [flags] <database.db>")
	}

	fileName := cmd.Args().Get(0)

	if !strings.HasSuffix(strings.ToLower(fileName), ".db") {
		return fmt.Errorf("invalid file '%s': extension must be .db", fileName)
	}

	return nil
}
