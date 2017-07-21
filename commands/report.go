package commands

type report struct {
}

func (r report) Type() commandType {
	return REPORT
}

func (r report) String() string {
	return "REPORT"
}
