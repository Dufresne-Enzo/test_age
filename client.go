package main

import (
	"fmt"
	"time"
)

type client struct {
	firstName    string
	lastName     string
	birthdayDate time.Time
	age          int
}

func newClient(lastName string, firstName string, birthdayDate string) *client {
	c := client{firstName, lastName, time.Time{}, 0}
	c.birthdayDate = parseDate(birthdayDate)
	c.age = getAge(c.birthdayDate.Year())
	c.age = dayYearCompare(c.birthdayDate.YearDay(), c.age)
	return &c
}

func getAge(year int) int {
	now := time.Now()
	return now.Year() - year
}

func parseDate(birthdayDate string) time.Time {
	layout := "01/02/06"
	t, _ := time.Parse(layout, birthdayDate)
	return t
}

func dayYearCompare(dayUser int, age int) int {
	now := time.Now()
	dayNow := now.YearDay()

	fmt.Println(dayUser)
	if dayUser >= dayNow {
		age--
	}
	return age
}
