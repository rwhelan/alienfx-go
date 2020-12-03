package zone

import "github.com/rwhelan/alienfx-go/internal"

const (
	R10_HEAD ZoneID = 0
	R10_RING ZoneID = 1
	R10_LOGO ZoneID = 3
)

type ZoneID uint8

type ZoneSelect []ZoneID

func (z *ZoneSelect) Serialize() []byte {
	zoneSelectHeader := []byte{0x03, 0x23, 0x01, 0x00}
	b := make([]byte, len(*z)+1)
	zids := []ZoneID(*z)

	b[0] = uint8(len(*z))

	for i := 0; i < len(*z); i++ {
		b[i+1] = byte(zids[i])
	}

	return internal.Pad(append(zoneSelectHeader, b...))
}

func (z *ZoneSelect) Add(zoneID ZoneID) {
	*z = append(*z, zoneID)
}
