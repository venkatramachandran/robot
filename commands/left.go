package commands

import "github.com/venkatramachandran/robot/objects"

type left struct {
}

func (l left) Execute(r *objects.Robot) {
	r.TurnLeft()
}

func (l left) String() string {
	return "LEFT"
}
