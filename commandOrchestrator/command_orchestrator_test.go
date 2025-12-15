package commandOrchestrator

// import (
// 	"database/sql"
// 	"testing"
// )

// func QueryFlagOrchestratorTest(t *testing.T) {
// 	// Load db
// 	db, err := sql.Open("sqlite3", "./test.db")

// 	if err != nil {
// 		t.Errorf("error: %v", err)
// 	}

// 	defer db.Close()

// 	result := queryFlagOrchestrator("SELECT * FROM test")

// 	if result != nil {
// 		t.Errorf("error: %v", result.Error())
// 	}
// }
