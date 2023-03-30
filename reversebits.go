package piscine

func ReverseBits(oct byte) byte {
	result := byte(0)
	for i := 0; i < 8; i++ {
		result <<= 1
		result = result | (oct & 1)
		oct >>= 1
	}
	return result
}
