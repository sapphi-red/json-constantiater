package lib

const hextable = "0123456789abcdef"

func AppendByteWithEscape(dest []byte, src string) []byte {
	for _, c := range []byte(src) {
		switch c {
		case '\\':
			dest = append(dest, '\\', '\\')
		case '"':
			dest = append(dest, '\\', '"')
		case '\n':
			dest = append(dest, '\\', 'n')
		case '\r':
			dest = append(dest, '\\', 'r')
		case '\t':
			dest = append(dest, '\\', 't')
		default:
			if c < byte(0x20) {
				dest = append(dest, 'u', '0', '0', hextable[c>>4], hextable[c&0x0f])
			} else {
				dest = append(dest, c)
			}
		}
	}
	return dest
}

func GetEscapedLen(src string) uint64 {
	var l uint64 = 0
	for _, c := range []byte(src) {
		switch c {
		case '\\', '"', '\n', '\r', '\t':
			l += 2
		default:
			if c < byte(0x20) {
				l += 5
			} else {
				l++
			}
		}
	}
	return l
}
