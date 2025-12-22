package query

import (
	"database/sql"
	"fmt"
)

func Execute(command string, query string, db *sql.DB) (*sql.Rows, error) {
	switch command {
	case "SELECT":
		rows, err := selectQuery(query, db)
		if err != nil {
			return nil, fmt.Errorf("query failed: %v", err)
		}
		return rows, nil

	case "INSERT", "CREATE", "UPDATE", "DELETE":
		if err := execQuery(query, db); err != nil {
			return nil, fmt.Errorf("query failed: %v", err)
		}
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid SQL query")
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
	return err
}
