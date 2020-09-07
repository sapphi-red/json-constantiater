package lib

const hextable = "0123456789abcdef"

func AppendString(dest []byte, src string) []byte {
	l := len(dest)
	dest = dest[:l+len(src)]
	copy(dest[l:], src)
	return dest
}

func AppendByteWithEscape(dest []byte, src string) []byte {
	p := len(dest)
	dest = dest[:p+len(src)]

	for _, c := range []byte(src) {
		switch c {
		case '\\':
			dest = dest[:p+1]
			dest[p] = '\\'
			dest[p+1] = '\\'
			p += 2
		case '"':
			dest = dest[:p+1]
			dest[p] = '\\'
			dest[p+1] = '"'
			p += 2
		case '\n':
			dest = dest[:p+1]
			dest[p] = '\\'
			dest[p+1] = 'n'
			p += 2
		case '\r':
			dest = dest[:p+1]
			dest[p] = '\\'
			dest[p+1] = 'r'
			p += 2
		case '\t':
			dest = dest[:p+1]
			dest[p] = '\\'
			dest[p+1] = 't'
			p += 2
		case byte(0x0), byte(0x1), byte(0x2), byte(0x3), byte(0x4), byte(0x5), byte(0x6), byte(0x7), byte(0x8), byte(0xb), byte(0xc), byte(0xe), byte(0xf), byte(0x10), byte(0x11), byte(0x12), byte(0x13), byte(0x14), byte(0x15), byte(0x16), byte(0x17), byte(0x18), byte(0x19), byte(0x1a), byte(0x1b), byte(0x1c), byte(0x1d), byte(0x1e), byte(0x1f):
			dest = dest[:p+4]
			dest[p] = 'u'
			dest[p+1] = '0'
			dest[p+2] = '0'
			dest[p+3] = hextable[c>>4]
			dest[p+4] = hextable[c&0x0f]
			p += 5
		default:
			dest[p] = c
			p++
		}
	}
	return dest
}

func GetEscapedLen(src string) uint64 {
	var l uint64 = uint64(len(src))
	for _, c := range []byte(src) {
		switch c {
		case '\\', '"', '\n', '\r', '\t':
			l ++
		case byte(0x0), byte(0x1), byte(0x2), byte(0x3), byte(0x4), byte(0x5), byte(0x6), byte(0x7), byte(0x8), byte(0xb), byte(0xc), byte(0xe), byte(0xf), byte(0x10), byte(0x11), byte(0x12), byte(0x13), byte(0x14), byte(0x15), byte(0x16), byte(0x17), byte(0x18), byte(0x19), byte(0x1a), byte(0x1b), byte(0x1c), byte(0x1d), byte(0x1e), byte(0x1f):
			l += 4
		}
	}
	return l
}
