package lib

const NSmalls = 1000

//go:nosplit
func AppendSmallInt(dest []byte, src *int) []byte {
	return appendSmallInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallInt8(dest []byte, src *int8) []byte {
	return appendSmallInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallInt16(dest []byte, src *int16) []byte {
	return appendSmallInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallInt32(dest []byte, src *int32) []byte {
	return appendSmallInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallInt64(dest []byte, src *int64) []byte {
	return appendSmallInt64(dest, int64(*src))
}

//go:nosplit
func appendSmallInt64(dest []byte, src int64) []byte {
	if src < 10 {
		return append(dest, digits[src])
	}
	if src < 100 {
		return append(dest, smallsString[src*2 : src*2+2]...)
	}
	return append(dest, smallsString2[src*3 : src*3+3]...)
}

//go:nosplit
func AppendSmallMinusInt(dest []byte, src *int) []byte {
	return appendSmallMinusInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallMinusInt8(dest []byte, src *int8) []byte {
	return appendSmallMinusInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallMinusInt16(dest []byte, src *int16) []byte {
	return appendSmallMinusInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallMinusInt32(dest []byte, src *int32) []byte {
	return appendSmallMinusInt64(dest, int64(*src))
}

//go:nosplit
func AppendSmallMinusInt64(dest []byte, src *int64) []byte {
	return appendSmallMinusInt64(dest, int64(*src))
}

//go:nosplit
func appendSmallMinusInt64(dest []byte, src int64) []byte {
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
func AppendSmallUint(dest []byte, src *uint) []byte {
	return appendSmallUint64(dest, uint64(*src))
}

//go:nosplit
func AppendSmallUint8(dest []byte, src *uint8) []byte {
	return appendSmallUint64(dest, uint64(*src))
}

//go:nosplit
func AppendSmallUint16(dest []byte, src *uint16) []byte {
	return appendSmallUint64(dest, uint64(*src))
}

//go:nosplit
func AppendSmallUint32(dest []byte, src *uint32) []byte {
	return appendSmallUint64(dest, uint64(*src))
}

//go:nosplit
func AppendSmallUint64(dest []byte, src *uint64) []byte {
	return appendSmallUint64(dest, uint64(*src))
}

//go:nosplit
func appendSmallUint64(dest []byte, i uint64) []byte {
	if i < 10 {
		return append(dest, digits[i])
	}
	if i < 100 {
		return append(dest, smallsString[i*2 : i*2+2]...)
	}
	return append(dest, smallsString2[i*3 : i*3+3]...)
}
