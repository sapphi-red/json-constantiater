package lib

import "unsafe"

//go:nosplit
func AppendByteWithEscape(dest []byte, src string) []byte {
	srcb := *(*[]byte)(unsafe.Pointer(&src))

	start := 0
	// avoid `i, c :=` for not copying `c`
	for i := range srcb {
		switch sizeTable[srcb[i]] {
		case 1:
			dest = append(dest, srcb[start:i]...)
			start = i + 1

			switch srcb[i] {
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
			dest = append(dest, srcb[start:i]...)
			start = i + 1

			dest = append(dest, 'u', '0', '0', hextable[srcb[i]>>4], hextable[srcb[i]&0x0f])
		}
	}

	if start < len(src) {
		dest = append(dest, srcb[start:]...)
	}
	return dest
}

//go:nosplit
func GetEscapedLen(src string) int {
	srcb := *(*[]byte)(unsafe.Pointer(&src))
	l := len(src)

	// avoid `i, c :=` for not copying `c`
	for i := range srcb {
		l += int(sizeTable[srcb[i]])
	}
	return l
}
