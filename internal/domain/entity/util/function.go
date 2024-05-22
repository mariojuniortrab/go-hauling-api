package util_entity

import (
	"fmt"
	"time"
)

func GetDateFromString(date string) (time.Time, error) {
	fmt.Println("GetDateFromString => chegou:", date)
	const shortForm = "2006-01-02"

	result, err := time.Parse(shortForm, date)
	fmt.Println("GetDateFromString => saiu:", result)

	return result, err
}

func GetStringFromDate(date time.Time) string {
	fmt.Println("GetStringFromDate => chegou:", date)
	const shortForm = "2006-01-02"
	fmt.Println("GetStringFromDate => saiu:", date.Format(shortForm))
	return date.Format(shortForm)
}
