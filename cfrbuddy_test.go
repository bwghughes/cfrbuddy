package cfrbuddy

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

var people []Person

func setup() {
	p := Person{
		Name:   "Ben",
		Number: "07899123456",
	}
	people = append(people, p)
}

func TestAddShiftWithOkTimings(t *testing.T) {
	is := is.New(t)
	shift, err := CreateShift(time.Now(), time.Now().Add(time.Hour*5), people)
	is.NoErr(err)
	is.True(shift.End.After(time.Now()))
}

func TestAddShiftLessThanAnHourFails(t *testing.T) {
	is := is.New(t)
	shift, err := CreateShift(time.Now(), time.Now().Add(time.Minute*59), people)
	is.True(fmt.Sprint(err) == "Shift can't be less than one hour")
	is.True(shift.Start.IsZero())
	is.True(shift.End.IsZero())
}

func TestAddShiftInThePastFails(t *testing.T) {
	is := is.New(t)
	shift, err := CreateShift(time.Now().Add(time.Hour*-5), time.Now(), people)
	is.True(fmt.Sprint(err) == "Shift start can't be in the past")
	is.True(shift.Start.IsZero())
	is.True(shift.End.IsZero())
}

func TestAddShiftWithNoPeopleFails(t *testing.T) {
	is := is.New(t)
	shift, err := CreateShift(time.Now(), time.Now().Add(time.Hour*5), []Person{})
	is.True(fmt.Sprint(err) == "Shift must have one or more people")
	is.True(shift.Start.IsZero())
	is.True(shift.End.IsZero())
}
