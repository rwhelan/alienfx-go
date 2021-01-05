package main

import (
	"fmt"
	"log"

	"github.com/google/gousb"
	"github.com/rwhelan/alienfx-go/pkg/instruction"
	"github.com/rwhelan/alienfx-go/pkg/zone"
)

var (
	transationStart  = []byte{0x03, 0x21, 0x00, 0x01, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	transationFinish = []byte{0x03, 0x21, 0x00, 0x03, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

func rainbow() [][]byte {
	zoneSelect := zone.ZoneSelect{
		zone.R10_HEAD,
		zone.R10_RING,
		zone.R10_LOGO,
	}

	instructions := instruction.InstructionSet{
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Red:      255,
		},
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Red:      255,
			Green:    255,
		},
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Green:    255,
		},
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Green:    255,
			Blue:     255,
		},
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Blue:     255,
		},
		instruction.Instruction{
			Type:     instruction.MORPH,
			Duration: 65535,
			Tempo:    1,
			Blue:     255,
			Red:      255,
		},
	}

	resp := [][]byte{
		transationStart,
		zoneSelect.Serialize(),
	}

	for _, inst := range instructions.Serialize() {
		resp = append(resp, inst)
	}

	return append(resp, transationFinish)
}

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(0x187c, 0x0550)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()

	if err = dev.SetAutoDetach(true); err != nil {
		log.Fatalf("Could not SetAutoDetach: %v", err)
	}

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalln("Unable to get config: ", err)
	}

	defer cfg.Close()

	for _, row := range rainbow() {
		i, err := dev.Control(
			0x21,  // REQUEST_TYPE_CLASS | RECIPIENT_INTERFACE | ENDPOINT_OUT
			9,     // SET_REPORT
			0x200, // "Vendor" Descriptor Type + 0 Descriptor Index
			0,     // USB interface 0
			row,
		)
		if err != nil {
			log.Fatalln("Err in ctrl:", err)
		}

		fmt.Println(i)
	}
}
