package commands

type right struct {
}

func (r right) Type() commandType {
	return RIGHT
}

func (r right) String() string {
	return "RIGHT"
}
