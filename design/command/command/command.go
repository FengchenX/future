package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}
func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}
func (c *StartCommand) Execute() {
	//c.mb.
}

type MotherBoard struct{}
func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}
