package lib

const NSmalls = 1000

//go:nosplit
func AppendSmallInt(dest []byte, src int) []byte {
	return AppendSmallInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallInt8(dest []byte, src int8) []byte {
	return AppendSmallInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallInt16(dest []byte, src int16) []byte {
	return AppendSmallInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallInt32(dest []byte, src int32) []byte {
	return AppendSmallInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallInt64(dest []byte, src int64) []byte {
	if src < 10 {
		return append(dest, digits[src])
	}
	if src < 100 {
		return append(dest, smallsString[src*2 : src*2+2]...)
	}
	return append(dest, smallsString2[src*3 : src*3+3]...)
}

//go:nosplit
func AppendSmallMinusInt(dest []byte, src int) []byte {
	return AppendSmallMinusInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallMinusInt8(dest []byte, src int8) []byte {
	return AppendSmallMinusInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallMinusInt16(dest []byte, src int16) []byte {
	return AppendSmallMinusInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallMinusInt32(dest []byte, src int32) []byte {
	return AppendSmallMinusInt64(dest, int64(src))
}

//go:nosplit
func AppendSmallMinusInt64(dest []byte, src int64) []byte {
	src = -src
	if src < 10 {
		return append(dest, '-', digits[src])
	}
	dest = append(dest, '-')
	if src < 100 {
		return append(dest, smallsString[src*2 : src*2+2]...)
	}
	return append(dest, smallsString2[src*3 : src*3+3]...)
}

//go:nosplit
func AppendSmallUint(dest []byte, src int) []byte {
	return AppendSmallUint64(dest, uint64(src))
}

//go:nosplit
func AppendSmallUint8(dest []byte, src int8) []byte {
	return AppendSmallUint64(dest, uint64(src))
}

//go:nosplit
func AppendSmallUint16(dest []byte, src int16) []byte {
	return AppendSmallUint64(dest, uint64(src))
}

//go:nosplit
func AppendSmallUint32(dest []byte, src int32) []byte {
	return AppendSmallUint64(dest, uint64(src))
}

//go:nosplit
func AppendSmallUint64(dest []byte, i uint64) []byte {
	if i < 10 {
		return append(dest, digits[i])
	}
	if i < 100 {
		return append(dest, smallsString[i*2 : i*2+2]...)
	}
	return append(dest, smallsString2[i*3 : i*3+3]...)
}
