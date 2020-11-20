package alienfx

import (
	"encoding/binary"
)

type InstrustionType uint8

const (
	COLOR_CMD InstrustionType = iota
	FLASH_CMD
	MORPH_CMD
)

type Instrustion struct {
	Type     InstrustionType
	Duration uint16
	Tempo    uint16
	Red      uint8
	Blue     uint8
	Green    uint8
}

func (inst *Instrustion) Serialize() *[8]byte {
	b := [8]byte{}

	b[0] = byte(inst.Type)
	b[5] = inst.Red
	b[6] = inst.Blue
	b[7] = inst.Green

	binary.BigEndian.PutUint16(b[1:3], inst.Duration)
	binary.BigEndian.PutUint16(b[3:5], inst.Tempo)

	return &b
}
