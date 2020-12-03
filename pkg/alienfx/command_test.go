package alienfx

import (
	"testing"

	"github.com/rwhelan/alienfx-go/pkg/instruction"
	"github.com/rwhelan/alienfx-go/pkg/zone"

	"github.com/google/go-cmp/cmp"
)

func TestCommandZoneBuilder(t *testing.T) {
	cmd := NewCommand()
	cmd.Zone(3).Zone(40).Zone(21)

	if cmd.zones[0] != 3 {
		t.Fatalf("%d != 3", cmd.zones[0])
	}

	if cmd.zones[1] != 40 {
		t.Fatalf("%d != 40", cmd.zones[1])
	}

	if cmd.zones[2] != 21 {
		t.Fatalf("%d != 21", cmd.zones[2])
	}
}

func TestCommandZoneMulti(t *testing.T) {
	cmd := NewCommand()
	cmd.Zone(8, 51, 32)

	if cmd.zones[0] != 8 {
		t.Fatalf("%d != 8", cmd.zones[0])
	}

	if cmd.zones[1] != 51 {
		t.Fatalf("%d != 51", cmd.zones[1])
	}

	if cmd.zones[2] != 32 {
		t.Fatalf("%d != 32", cmd.zones[2])
	}
}

func TestCommandSerialize(t *testing.T) {
	tests := []struct {
		name   string
		zones  []zone.ZoneID
		is     []instruction.Instruction
		result [][]byte
	}{
		{
			name: "Seven Instructions",
			zones: []zone.ZoneID{
				18, 27, 31,
			},
			is: []instruction.Instruction{
				instruction.Instruction{
					Type:     instruction.MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.COLOR,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				instruction.Instruction{
					Type:     instruction.FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
			},
			result: [][]byte{
				[]byte{
					0x03, 0x23, 0x01, 0x00, // Zone Select Header
					0x03,             // Number of zones
					0x12, 0x1b, 0x1f, // Zone numbers
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00,
				},
				[]byte{
					0x03, 0x24, // Header
					0x02, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // First Instruction
					0x02, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Second Instruction
					0x00, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Third Instruction
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
				},
				[]byte{
					0x03, 0x24, // Header
					0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Forth Instruction
					0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Fifth Instruction
					0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Sixth Instruction
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
				},
				[]byte{
					0x03, 0x24, // Header
					0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // Seventh Instruction
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding ...
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand()
			for _, z := range tt.zones {
				cmd.Zone(z)
			}

			for _, i := range tt.is {
				cmd.Instruction(i)
			}

			b := cmd.Serialize()
			if diff := cmp.Diff(tt.result, b); diff != "" {
				t.Fatalf("Failed test \"%s\" (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
