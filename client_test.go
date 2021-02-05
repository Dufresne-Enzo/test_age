package main

import (
	"testing"
)

func testParseDate(t *testing.T) {
	var birthdayDate string

	birthdayDate = "01/06/1997"

	date := parseDate(birthdayDate)

	if date == 0 {
		t.Error("La date est egal à 0")
	}
	if date != 1997{
		t.Error("L'année ne correspond pas à la date d'anniversaire")
	}
}

func testGetAge(t *testing.T) {

	var y int
	y = 1997
	age := getAge(y)
	if age == 0 {
		t.Error("L age est égal à 0")
	}
	if age <= 0 {
		t.Error("L age est négatif")
	}
}