package commands

import (
	"strconv"
	"strings"

	"github.com/venkatramachandran/robot/errors"
	"github.com/venkatramachandran/robot/models"
	"github.com/venkatramachandran/robot/objects"
)

//Command is a command that a robot can understand
type Command interface {
	//Execute executes the command on the robot
	Execute(r *objects.Robot)
}

//From creates a Command object from a string
func From(s string) (Command, error) {
	if s == "LEFT" {
		return left{}, nil
	}
	if s == "RIGHT" {
		return right{}, nil
	}
	if s == "MOVE" {
		return move{}, nil
	}
	if s == "REPORT" {
		return report{}, nil
	}
	if strings.HasPrefix(s, "PLACE ") {
		posAndDir := strings.Split(s[6:], ",")
		if len(posAndDir) != 3 {
			return nil, errors.InvalidCommandError{Command: s}
		}
		x, err := strconv.Atoi(strings.TrimSpace(posAndDir[0]))
		if err != nil {
			return nil, errors.InvalidCommandError{Command: s}
		}
		y, err := strconv.Atoi(strings.TrimSpace(posAndDir[1]))
		if err != nil {
			return nil, errors.InvalidCommandError{Command: s}
		}
		dir, err := models.DirectionFrom(strings.TrimSpace(posAndDir[2]))
		if err != nil {
			return nil, errors.InvalidCommandError{Command: s}
		}
		return place{Position: models.Position{X: x, Y: y}, Direction: dir}, nil
	}
	return nil, errors.InvalidCommandError{Command: s}
}
