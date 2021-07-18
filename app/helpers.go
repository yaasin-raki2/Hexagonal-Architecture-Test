package app

import "log"

func LogIfErr(err error) {
	if err != nil {
		log.Printf("error occurred: %s", err)
	}
}
