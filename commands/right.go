package commands

import (
	"github.com/venkatramachandran/robot/objects"
)

type right struct {
}

func (r right) Execute(robot *objects.Robot) {
	robot.TurnRight()
}

func (r right) String() string {
	return "RIGHT"
}
