package lib

import (
	_ "unsafe"
)

//go:nosplit
func AppendInt(dest []byte, src int) []byte {
	return AppendInt64(dest, int64(src))
}

//go:nosplit
func AppendInt16(dest []byte, src int16) []byte {
	return AppendInt64(dest, int64(src))
}

//go:nosplit
func AppendInt32(dest []byte, src int32) []byte {
	return AppendInt64(dest, int64(src))
}

//go:nosplit
func AppendInt64(dest []byte, src int64) []byte {
	return formatBits(dest, uint64(src), src < 0)
}

//go:nosplit
func AppendUint(dest []byte, src uint) []byte {
	return AppendUint64(dest, uint64(src))
}

//go:nosplit
func AppendUint16(dest []byte, src uint16) []byte {
	return AppendUint64(dest, uint64(src))
}

//go:nosplit
func AppendUint32(dest []byte, src uint32) []byte {
	return AppendUint64(dest, uint64(src))
}

//go:nosplit
func AppendUint64(dest []byte, src uint64) []byte {
	return formatBits(dest, src, true)
}

//go:nosplit
func formatBits(dst []byte, u uint64, neg bool) []byte {
	d, _ := strconv_formatBits(dst, u, 10, neg, true)
	return d
}

//go:linkname strconv_formatBits strconv.formatBits
func strconv_formatBits(dst []byte, u uint64, base int, neg, append_ bool) (d []byte, s string)
