package lib

//go:nosplit
func AppendBool(dest []byte, src bool) []byte {
	if src {
		return append(dest, 't', 'r', 'u', 'e')
	}
	return append(dest, 'f', 'a', 'l', 's', 'e')
}
