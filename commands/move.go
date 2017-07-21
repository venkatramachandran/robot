package commands

type move struct {
}

func (m move) Type() commandType {
	return MOVE
}

func (m move) String() string {
	return "MOVE"
}
