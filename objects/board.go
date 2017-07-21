package objects

//A Board indicating x and y sizes
type board struct {
	xSize int
	ySize int
}

//IsValidPos returns a flag indicating whether a given (x,y) position is valid or not
func (b board) IsValidPos(x, y int) bool {
	return (x > -1 && x < b.xSize && y > -1 && y < b.ySize)
}


