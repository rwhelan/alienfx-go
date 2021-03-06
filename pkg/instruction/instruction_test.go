package instruction

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInstructionSerialize(t *testing.T) {
	tests := []struct {
		name string
		m    *Instruction
		b    []byte
	}{
		{
			name: "One",
			m: &Instruction{
				Type:     FLASH,
				Duration: 16328,
				Tempo:    253,
				Red:      251,
				Blue:     231,
				Green:    128,
			},
			b: []byte{
				0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.m.Serialize()
			if diff := cmp.Diff(tt.b, b); diff != "" {
				t.Fatalf("Failed test \"%s\" (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}

func TestInstructionSet(t *testing.T) {
	tests := []struct {
		name string
		m    *InstructionSet
		b    [][]byte
	}{
		{
			name: "One Instruction",
			m: &InstructionSet{
				Instruction{
					Type:     MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
			},
			b: [][]byte{
				[]byte{
					0x03, 0x24, // Header
					0x02, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0x80, 0xe7, // First Instruction
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			name: "Four Instructions",
			m: &InstructionSet{
				Instruction{
					Type:     MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     COLOR,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
			},
			b: [][]byte{
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
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding ...
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			name: "Seven Instructions",
			m: &InstructionSet{
				Instruction{
					Type:     MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     MORPH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     COLOR,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
				Instruction{
					Type:     FLASH,
					Duration: 16328,
					Tempo:    253,
					Red:      251,
					Blue:     231,
					Green:    128,
				},
			},
			b: [][]byte{
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
			b := tt.m.Serialize()
			if diff := cmp.Diff(tt.b, b); diff != "" {
				t.Fatalf("Failed test \"%s\" (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
