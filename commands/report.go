package commands

import "github.com/venkatramachandran/robot/objects"

type report struct {
}

func (r report) Execute(robot *objects.Robot) {
	robot.Report()
}

func (r report) String() string {
	return "REPORT"
}
