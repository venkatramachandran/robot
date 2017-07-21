package models

import "github.com/venkatramachandran/robot/errors"

//Direction is an enum for the four directions
type Direction int

//Direction constants
const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

func (d Direction) String() string {
	switch d {
	case NORTH:
		return "NORTH"
	case SOUTH:
		return "SOUTH"
	case EAST:
		return "EAST"
	case WEST:
		return "WEST"
	}
	return ""
}

//DirectionFrom returns a direction constant from a string input
func DirectionFrom(s string) (Direction, error) {
	if s == "EAST" {
		return EAST, nil
	}
	if s == "WEST" {
		return WEST, nil
	}
	if s == "NORTH" {
		return NORTH, nil
	}
	if s == "SOUTH" {
		return SOUTH, nil
	}
	return -1, errors.InvalidDirectionError{Direction: s}
}
