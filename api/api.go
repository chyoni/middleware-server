package api

import (
	"database/sql"
)

func CountRecords(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT COUNT(*) FROM message")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}
	return count, nil
}
