package errors

import "fmt"

//InvalidCommandError indicates when a string input cannot be converted to a command object
type InvalidCommandError struct {
	Command string
}

func (i InvalidCommandError) Error() string {
	return fmt.Sprintf("Invalid Command: %s", i.Command)
}

//InvalidDirectionError indicates when a string input cannot be converted to a direction object
type InvalidDirectionError struct {
	Direction string
}

func (i InvalidDirectionError) Error() string {
	return fmt.Sprintf("Invalid Direction: %s", i.Direction)
}
