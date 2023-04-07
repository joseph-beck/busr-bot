package util

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrMsg(err error, msg string) {
	if err != nil {
		log.Fatal(msg+" ", err)
	}
}
