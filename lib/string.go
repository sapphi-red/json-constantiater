package lib

//go:nosplit
func AppendByteWithEscape(dest []byte, src string) []byte {
	for _, c := range []byte(src) {
		switch sizeTable[c] {
		case 0:
			dest = append(dest, c)
		case 1:
			switch c {
			case '"':
				dest = append(dest, "\\\""...)
			case '\\':
				dest = append(dest, "\\\\"...)
			case '\n':
				dest = append(dest, "\\n"...)
			case '\r':
				dest = append(dest, "\\r"...)
			case '\t':
				dest = append(dest, "\\t"...)
			}
		case 4:
			dest = append(dest, 'u', '0', '0', hextable[c>>4], hextable[c&0x0f])
		}
	}
	return dest
}

//go:nosplit
func GetEscapedLen(src string) uint64 {
	var l uint64 = uint64(len(src))
	for _, c := range []byte(src) {
		l += uint64(sizeTable[c])
	}
	return l
}
