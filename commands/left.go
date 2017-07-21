package commands

type left struct {
}

func (l left) Type() commandType {
	return LEFT
}

func (l left) String() string {
	return "LEFT"
}
