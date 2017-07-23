package models

import (
	"testing"

	"github.com/venkatramachandran/robot/errors"
)

type testCase struct {
	Input string
	Direction
	Error error
}

var testData = []testCase{
	{"EAST", EAST, nil},
	{"WEST", WEST, nil},
	{"NORTH", NORTH, nil},
	{"SOUTH", SOUTH, nil},
	{"LEFT", -1, errors.InvalidDirectionError{Direction: "LEFT"}},
}

func TestParseDirection(t *testing.T) {
	for i, data := range testData {
		dir, err := DirectionFrom(data.Input)
		if dir != data.Direction || err != data.Error {
			t.Fatalf("testcase %d failed! Expected (command,error): (%v, %s), got (%v, %s)\n",
				i, data.Direction, data.Error, dir, err)
		}
	}

}
