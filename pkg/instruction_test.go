package alienfx

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
				0x01, 0x3f, 0xc8, 0x00, 0xfd, 0xfb, 0xe7, 0x80,
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
