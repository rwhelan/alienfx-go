package msgtypes

func pad(b []byte) []byte {
	return append(b, make([]byte, 33-len(b))...)
}
