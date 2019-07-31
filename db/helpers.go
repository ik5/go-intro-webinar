package db

import "database/sql"

// IsNoRows return true if a given err is about no rows were returned
func IsNoRows(err error) bool {
	return err == sql.ErrNoRows
}
