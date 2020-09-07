package lib

func AppendBool(dest []byte, src bool) []byte {
	if src {
		return append(dest, "true"...)
	}
	return append(dest, "false"...)
}
