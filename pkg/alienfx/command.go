package alienfx

import (
	"github.com/rwhelan/alienfx-go/pkg/instruction"
	"github.com/rwhelan/alienfx-go/pkg/zone"
)

type Command struct {
	zones        zone.ZoneSelect
	instructions instruction.InstructionSet
}

func (c *Command) Serialize() [][]byte {
	insts := c.instructions.Serialize()

	resp := make([][]byte, 0, len(insts)+1)

	resp = append(resp, c.zones.Serialize())
	resp = append(resp, insts...)

	return resp
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Zone(zoneID ...zone.ZoneID) *Command {
	for _, z := range zoneID {
		c.zones.Add(z)
	}
	return c
}

func (c *Command) Instruction(instruction ...instruction.Instruction) *Command {
	for _, i := range instruction {
		c.instructions.Add(i)
	}
	return c
}
