package lib

const hextable = "0123456789abcdef"

func AppendByteWithEscape(dest []byte, src string) []byte {
	for _, c := range []byte(src) {
		if c == '"' || c == '\\' {
			dest = append(dest, '\\', c)
		} else if c < 0x20 {
			if c == '\n' {
				dest = append(dest, '\\', 'n')
			} else if c == '\r' {
				dest = append(dest, '\\', 'r')
			} else if c == '\t' {
				dest = append(dest, '\\', 't')
			} else {
				dest = append(dest, 'u', '0', '0', hextable[c>>4], hextable[c&0x0f])
			}
		} else {
			dest = append(dest, c)
		}
	}
	return dest
}

func GetEscapedLen(src string) uint64 {
	var l uint64 = uint64(len(src))
	for _, c := range []byte(src) {
		if c == '"' || c == '\\' {
			l++
		} else if c < 0x20 {
			if c == '\n' || c == '\r' || c == '\t' {
				l++
			} else {
				l += 4
			}
		}
	}
	return l
}
