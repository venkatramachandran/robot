package commands

import (
	"fmt"

	"github.com/venkatramachandran/robot/models"
	"github.com/venkatramachandran/robot/objects"
)

//Place is a command which places the robot in a location and sets its direction
type place struct {
	models.Position
	models.Direction
}

func (p place) Execute(r *objects.Robot) {
	r.SetPosition(p.Position, p.Direction)
}

func (p place) String() string {
	return fmt.Sprintf("PLACE %d, %d, %s", p.Position.X, p.Position.Y, p.Direction.String())
}
