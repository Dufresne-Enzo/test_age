package main

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	var birthdayDate string

	birthdayDate = "01/06/97"

	date := parseDate(birthdayDate)

	if date == 0 {
		t.Error("La date est egal à 0")
	}
	if date != 1997 {
		t.Error("L'année ne correspond pas à la date d'anniversaire")
	}
}

func TestGetAge(t *testing.T) {

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

func TestCompareDays(t *testing.T) {
	dateUser := time.Date(1997, 6, 1, 0, 0, 0, 0, time.UTC)
	dayUser := dateUser.YearDay()

	age := 24
	age = dayYearCompare(dayUser, age)

	if age >= 24 {
		t.Error("Age incorrect ")
	}
}
