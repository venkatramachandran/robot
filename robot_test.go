package main

import (
	"testing"

	"github.com/venkatramachandran/robot/commands"
	"github.com/venkatramachandran/robot/models"
	"github.com/venkatramachandran/robot/objects"
)

type testCase struct {
	Input        []string
	EndPosition  models.Position
	EndDirection models.Direction
}

var testData = []testCase{
	{[]string{"PLACE 0,0,EAST", "MOVE", "LEFT", "MOVE", "MOVE", "LEFT"}, models.Position{X: 1, Y: 2}, models.WEST},
	{[]string{"PLACE 0,0,EAST", "RIGHT", "RIGHT"}, models.Position{X: 0, Y: 0}, models.WEST},
	{[]string{"MOVE", "MOVE", "LEFT", "PLACE 0,0,WEST", "LEFT", "LEFT"}, models.Position{X: 0, Y: 0}, models.EAST},
}

func TestRobot(t *testing.T) {
	for i, data := range testData {
		r := objects.NewRobot(5, 5)
		for _, c := range data.Input {
			cmd, _ := commands.From(c)
			r.Execute(cmd)
		}
		if r.Position.X != data.EndPosition.X || r.Position.Y != data.EndPosition.Y || r.Direction != data.EndDirection {
			t.Fatalf("test case %d failed! Expected Position: %d, %d, %s. Got %d, %d, %s", i,
				data.EndPosition.X, data.EndPosition.Y, data.EndDirection.String(),
				r.Position.X, r.Position.Y, r.Direction.String())
		}
	}
}
