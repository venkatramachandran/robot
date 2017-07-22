package commands

import "github.com/venkatramachandran/robot/objects"

type move struct {
}

func (m move) Execute(r *objects.Robot) {
	r.MoveForward()
}

func (m move) String() string {
	return "MOVE"
}
