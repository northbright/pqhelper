package pqhelper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// TableExists checks if given table exists in given scheme.
func TableExists(db *sql.DB, schemaName, dbName string) (exists bool, err error) {
	exists = false
	result := ""
	sqlStat := `
    SELECT EXISTS (
        SELECT 1
        FROM pg_tables
        WHERE schemaname = $1
        AND tablename = $2
    );`

	if err = db.QueryRow(sqlStat, schemaName, dbName).Scan(&result); err != nil {
		fmt.Printf("db.QueryRow() error: %v\n", err)
		goto end
	}

	if result == "true" {
		exists = true
	}

end:
	return exists, err
}
