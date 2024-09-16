package main

import "time"

type Month struct {
	title     string
	numOfDays int
	year      int
}

type Year struct {
	year      int
	January   Month
	February  Month
	March     Month
	April     Month
	May       Month
	June      Month
	July      Month
	August    Month
	September Month
	October   Month
	November  Month
	December  Month
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

func createYear(year int) Year {
	febDays := 28
	if isLeapYear(year) {
		febDays = 29
	}

	return Year{
		year:      year,
		January:   Month{title: "January", numOfDays: 31, year: year},
		February:  Month{title: "February", numOfDays: febDays, year: year},
		March:     Month{title: "March", numOfDays: 31, year: year},
		April:     Month{title: "April", numOfDays: 30, year: year},
		May:       Month{title: "May", numOfDays: 31, year: year},
		June:      Month{title: "June", numOfDays: 30, year: year},
		July:      Month{title: "July", numOfDays: 31, year: year},
		August:    Month{title: "August", numOfDays: 31, year: year},
		September: Month{title: "September", numOfDays: 30, year: year},
		October:   Month{title: "October", numOfDays: 31, year: year},
		November:  Month{title: "November", numOfDays: 30, year: year},
		December:  Month{title: "December", numOfDays: 31, year: year},
	}
}

func (y Year) daysInMonth(m Month) int {
	return m.numOfDays
}

func (y Year) createDaysOfYear() []Date {
	days := []Date{}
	months := []Month{y.January, y.February, y.March, y.April, y.May, y.June, y.July, y.August, y.September, y.October, y.November, y.December}
	for _, month := range months {
		for i := 1; i <= month.numOfDays; i++ {
			days = append(days, Date{
				dayOfWeek: time.Date(y.year, time.Month(i), i, 0, 0, 0, 0, time.UTC).Weekday(),
				date:      i,
				month:     time.Date(y.year, time.Month(i), i, 0, 0, 0, 0, time.UTC).Month(),
				year:      y.year,
			})
		}
	}
	return days
}
