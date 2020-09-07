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
		case byte(0x0), byte(0x1), byte(0x2), byte(0x3), byte(0x4), byte(0x5), byte(0x6), byte(0x7), byte(0x8), byte(0xb), byte(0xc), byte(0xe), byte(0xf), byte(0x10), byte(0x11), byte(0x12), byte(0x13), byte(0x14), byte(0x15), byte(0x16), byte(0x17), byte(0x18), byte(0x19), byte(0x1a), byte(0x1b), byte(0x1c), byte(0x1d), byte(0x1e), byte(0x1f):
			dest = append(dest, 'u', '0', '0', hextable[c>>4], hextable[c&0x0f])
		default:
			dest = append(dest, c)
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
		case byte(0x0), byte(0x1), byte(0x2), byte(0x3), byte(0x4), byte(0x5), byte(0x6), byte(0x7), byte(0x8), byte(0xb), byte(0xc), byte(0xe), byte(0xf), byte(0x10), byte(0x11), byte(0x12), byte(0x13), byte(0x14), byte(0x15), byte(0x16), byte(0x17), byte(0x18), byte(0x19), byte(0x1a), byte(0x1b), byte(0x1c), byte(0x1d), byte(0x1e), byte(0x1f):
			l += 5
		default:
			l++
		}
	}
	return l
}
