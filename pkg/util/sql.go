package util

import (
	"database/sql"
	"log"
)

func CheckErrMsg(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func CheckRow(err error) (bool, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
