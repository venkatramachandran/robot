package commands

import "fmt"
import "github.com/venkatramachandran/robot/models"

//Place is a command which places the robot in a location and sets its direction
type Place struct {
	models.Position
	models.Direction
}

//Type returns the command type
func (p Place) Type() commandType {
	return PLACE
}

func (p Place) String() string {
	return fmt.Sprintf("PLACE %d, %d, %s", p.Position.X, p.Position.Y, p.Direction.String())
}
