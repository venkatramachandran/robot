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
	//place first and move
	{[]string{"PLACE 0,0,EAST", "MOVE", "LEFT", "MOVE", "MOVE", "LEFT"}, models.Position{X: 1, Y: 2}, models.WEST},
	//place first and just turn
	{[]string{"PLACE 0,0,EAST", "RIGHT", "RIGHT"}, models.Position{X: 0, Y: 0}, models.WEST},
	//try moves before placing
	{[]string{"MOVE", "MOVE", "LEFT", "PLACE 0,0,WEST", "LEFT", "LEFT"}, models.Position{X: 0, Y: 0}, models.EAST},
	//sample input 1 and output given in the problem
	{[]string{"PLACE 0, 0, NORTH", "MOVE"}, models.Position{X: 0, Y: 1}, models.NORTH},
	//sample input 2 and output given in the problem
	{[]string{"PLACE 0, 0, NORTH", "LEFT"}, models.Position{X: 0, Y: 0}, models.WEST},
	//sample input 3 and output given in the problem
	{[]string{"PLACE 1, 2, EAST", "MOVE", "MOVE", "LEFT", "MOVE"}, models.Position{X: 3, Y: 3}, models.NORTH},
	//make sure the robot does not fall down
	{[]string{"PLACE 3,2, EAST", "MOVE", "MOVE", "MOVE", "LEFT", "MOVE"}, models.Position{X: 4, Y: 3}, models.NORTH},
	//another test to make sure the robot does not fall down
	{[]string{"PLACE 0,0, NORTH", "MOVE", "MOVE", "MOVE", "MOVE", "RIGHT", "MOVE", "MOVE", "MOVE", "MOVE", "MOVE", "LEFT"}, models.Position{X: 4, Y: 4}, models.NORTH},
}

func TestRobot(t *testing.T) {
	for i, data := range testData {
		r := objects.NewRobot(5, 5)
		for _, c := range data.Input {
			cmd, _ := commands.From(c)
			cmd.Execute(r)
		}
		if r.Position.X != data.EndPosition.X || r.Position.Y != data.EndPosition.Y || r.Direction != data.EndDirection {
			t.Fatalf("test case %d failed! Expected Position: %d, %d, %s. Got %d, %d, %s", i,
				data.EndPosition.X, data.EndPosition.Y, data.EndDirection.String(),
				r.Position.X, r.Position.Y, r.Direction.String())
		}
	}
}
