package utils

// BytesToInt Bytes数组转Int
func BytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return result
}

// IntToBytes Int转Bytes数组
func IntToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}

func IntToString(n []int) string{
	result := ""
	for _,v := range n{
		result += string(v)
	}
	return result
}
