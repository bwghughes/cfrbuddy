package cfrbuddy

import (
	"errors"
	"time"
)

type Person struct {
	Name   string
	Number string
}

type Shift struct {
	Start  time.Time
	End    time.Time
	People []Person
}

func CreateShift(start time.Time, end time.Time, people []Person) (shift Shift, err error) {

	// Shift time window and people validation
	if time.Since(start).Milliseconds() > 0 {
		return Shift{}, errors.New("Shift start can't be in the past")
	} else if time.Since(end).Minutes() > -60 {
		return Shift{}, errors.New("Shift can't be less than one hour")
	} else if len(people) == 0 {
		return Shift{}, errors.New("Shift must have one or more people")
	} else {
		return Shift{Start: start, End: end, People: people}, nil
	}

}
