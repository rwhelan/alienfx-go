package msgtypes

import (
	"encoding/binary"
)

type InstructionType uint8

const (
	COLOR InstructionType = iota
	FLASH
	MORPH
)

type Instruction struct {
	Type     InstructionType
	Duration uint16
	Tempo    uint16
	Red      uint8
	Blue     uint8
	Green    uint8
}

func (inst *Instruction) Serialize() []byte {
	b := [8]byte{}

	b[0] = byte(inst.Type)
	b[5] = inst.Red
	b[6] = inst.Blue
	b[7] = inst.Green

	binary.BigEndian.PutUint16(b[1:3], inst.Duration)
	binary.BigEndian.PutUint16(b[3:5], inst.Tempo)

	return b[:]
}

type InstructionSet []Instruction

func (instSet *InstructionSet) Serialize() [][]byte {
	instHeader := []byte{0x03, 0x24}

	count := 3 - (len(*instSet) % 3)
	batches := make([][]Instruction, 0, count)
	b := make([][]byte, count)

	for i := 0; i < len(*instSet); i += 3 {
		j := i + 3
		if j > len(*instSet) {
			j = len(*instSet)
		}

		batches = append(batches, []Instruction(*instSet)[i:j])
	}

	for i, set := range batches {
		for _, subset := range set {
			b[i] = append(b[i], subset.Serialize()...)
		}
	}

	for i := range b {
		b[i] = pad(append(instHeader, b[i]...))
	}

	return b
}
