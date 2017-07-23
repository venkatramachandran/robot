package commands

import (
	"testing"

	"github.com/venkatramachandran/robot/errors"
	"github.com/venkatramachandran/robot/models"
)

type testCase struct {
	Input string
	Command
	Error error
}

var testData = []testCase{
	{"MOVE", move{}, nil},
	{"LEFT", left{}, nil},
	{"RIGHT", right{}, nil},
	{"LEFT", left{}, nil},
	{"REPORT", report{}, nil},
	{"PLACE 1, 2, EAST", place{Position: models.Position{X: 1, Y: 2}, Direction: models.EAST}, nil},
	{"BLAH", nil, errors.InvalidCommandError{Command: "BLAH"}},
	{"PLACE 1,2, NORTHWEST", nil, errors.InvalidDirectionError{Direction: "NORTHWEST"}},
}

func TestParseCommand(t *testing.T) {
	for i, data := range testData {
		cmd, err := From(data.Input)
		if cmd != data.Command || err != data.Error {
			t.Fatalf("testcase %d failed! Expected (command,error): (%v, %s), got (%v, %s)\n",
				i, data.Command, data.Error, cmd, err)
		}
	}
}
