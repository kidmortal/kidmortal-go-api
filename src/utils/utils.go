package utils

import (
	"log"
	"time"
)

func DataStringToTime(Data string) time.Time {

	data, err := time.Parse("02/01/2006", Data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
