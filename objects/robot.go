package objects

import (
	"fmt"
	"strconv"

	"github.com/venkatramachandran/robot/commands"
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

//Execute a command object on the robot
func (r *Robot) Execute(c commands.Command) {
	switch c.Type() {
	case commands.REPORT:
		r.report()
	case commands.MOVE:
		r.moveForward()
	case commands.LEFT:
		r.turnLeft()
	case commands.RIGHT:
		r.turnRight()
	case commands.PLACE:
		if p, ok := c.(commands.Place); ok {
			r.setPosition(p.Position, p.Direction)
		}
	}
}

//NewRobot creates a new Robot
func NewRobot(xSize, ySize int) *Robot {
	return &Robot{Position: models.Position{X: -1, Y: -1},
		Direction: -1,
		board:     board{xSize: xSize, ySize: ySize}}
}

func (r *Robot) moveForward() {
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

func (r *Robot) turnLeft() {
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

func (r *Robot) turnRight() {
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

func (r *Robot) setPosition(p models.Position, d models.Direction) {
	if r.board.IsValidPos(p.X, p.X) {
		r.Position.X = p.X
		r.Position.Y = p.Y
		r.Direction = d
	}
}

func (r Robot) report() {
	if !r.isPlaced() {
		return
	}

	fmt.Printf("%d, %d, %s\n", r.Position.X, r.Position.Y, r.Direction)
}

func (r Robot) String() string {
	return strconv.Itoa(r.Position.X) + "," + strconv.Itoa(r.Position.Y) + "," + r.Direction.String()
}
