package main

import "time"

type Date struct {
	dayOfWeek time.Weekday
	date      int
	month     time.Month
	year      int
}

func dateFromTime(t time.Time) Date {
	return Date{
		dayOfWeek: t.Weekday(),
		date:      t.Day(),
		month:     t.Month(),
		year:      t.Year(),
	}
}

func (d Date) timeFromDate() time.Time {
	return time.Date(d.year, d.month, d.date, 0, 0, 0, 0, time.UTC)
}

func newTime(t time.Time, year int, month int, days int) time.Time {
	return t.AddDate(year, month, days)
}
