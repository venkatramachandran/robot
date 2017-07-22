package objects

import (
	"fmt"
	"strconv"

	"github.com/venkatramachandran/robot/models"
)

//A Robot and its board
type Robot struct {
	models.Position
	models.Direction
	board
}

func (r Robot) isPlaced() bool {
	return !(r.Position.X == -1 || r.Position.Y == -1)
}

//NewRobot creates a new Robot
func NewRobot(xSize, ySize int) *Robot {
	return &Robot{Position: models.Position{X: -1, Y: -1},
		Direction: -1,
		board:     board{xSize: xSize, ySize: ySize}}
}

//MoveForward moves the robot one step forward in the direction it is facing
func (r *Robot) MoveForward() {
	//move forward one step in the current direction
	//if it makes the robot exceed the limits, ignore
	if !r.isPlaced() {
		return
	}

	switch r.Direction {
	case models.NORTH:
		if r.board.IsValidPos(r.Position.X, r.Position.Y+1) {
			r.Position.Y = r.Position.Y + 1
		}
	case models.SOUTH:
		if r.board.IsValidPos(r.Position.X, r.Position.Y-1) {
			r.Position.Y = r.Position.Y - 1
		}
	case models.EAST:
		if r.board.IsValidPos(r.Position.X+1, r.Position.Y) {
			r.Position.X = r.Position.X + 1
		}
	case models.WEST:
		if r.board.IsValidPos(r.Position.X-1, r.Position.Y) {
			r.Position.X = r.Position.X - 1
		}
	}
}

//TurnLeft makes the robot turn to its left
func (r *Robot) TurnLeft() {
	//move 90 degrees to the left
	if !r.isPlaced() {
		return
	}

	switch r.Direction {
	case models.NORTH:
		r.Direction = models.WEST
	case models.SOUTH:
		r.Direction = models.EAST
	case models.EAST:
		r.Direction = models.NORTH
	case models.WEST:
		r.Direction = models.SOUTH
	}
}

//TurnRight makes the robot turn to its right
func (r *Robot) TurnRight() {
	//move 90 degrees to the left

	if !r.isPlaced() {
		return
	}
	switch r.Direction {
	case models.NORTH:
		r.Direction = models.EAST
	case models.SOUTH:
		r.Direction = models.WEST
	case models.EAST:
		r.Direction = models.SOUTH
	case models.WEST:
		r.Direction = models.NORTH
	}
}

//SetPosition sets the initial position and direction for the robot
func (r *Robot) SetPosition(p models.Position, d models.Direction) {
	if r.board.IsValidPos(p.X, p.X) {
		r.Position.X = p.X
		r.Position.Y = p.Y
		r.Direction = d
	}
}

//Report prints the current position and direction of the robot
func (r Robot) Report() {
	if !r.isPlaced() {
		return
	}

	fmt.Printf("%d, %d, %s\n", r.Position.X, r.Position.Y, r.Direction)
}

func (r Robot) String() string {
	return strconv.Itoa(r.Position.X) + "," + strconv.Itoa(r.Position.Y) + "," + r.Direction.String()
}
