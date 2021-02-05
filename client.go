package main

import (
	"time"
	"fmt"
)

type client struct {
    firstName string
	lastName  string
	birthdayDate string
	age int
}

func newClient(lastName string, firstName string, birthdayDate string) *client {
	c := client{firstName, lastName, birthdayDate, 0}
	year := parseDate(c.birthdayDate)
	c.age = getAge(year)
    return &c
}

func getAge(year int) int {
	now := time.Now()
	return now.Year() - year
}

func parseDate(birthdayDate string) int {
	layout := "01/02/06"
	t, _ := time.Parse(layout, birthdayDate)
	return t.Year()
}
	