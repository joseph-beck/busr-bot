package util

import (
	"database/sql"
)

func CheckRow(err error) (bool, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
