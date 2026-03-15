package database

import (
	"database/sql"
	"fmt"
)

func GetListFromRows[K string | int](rows *sql.Rows, err error) ([]K, error) {
	var arr []K
	if err == nil {
		for rows.Next() {
			var v K
			err = rows.Scan(&v)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return arr, fmt.Errorf("GetListFromRows: %v", err)
			}
			arr = append(arr, v)
		}
	}
	if rows != nil {
		rows.Close()
	}
	return arr, nil
}
