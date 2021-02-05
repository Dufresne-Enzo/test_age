package client

import (
	"testing"
	"time"
	"fmt"
)

func testParseDate(t *testing.T) {
	var birthdayDate string

	birthdayDate = "01/06/1997"

	date := parseDate(birthdayDate)

	if date == nil {
		t.Error("La date est null")
	}
	if date.Year() != 1997{
		t.Error("L'année ne correspond pas à la date d'anniversaire")
	}
}

func testGetAge(t *testing.T) {

	var y int
	y = 1997
	age := getAge(y)
	if age == nil {
		t.Error("L age est null")
	}
	if age <= 0 {
		t.Error("L age est négatif")
	}
}