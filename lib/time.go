package lib

import "time"

//go:nosplit
func AppendTime(dest []byte, src *time.Time) []byte {
	dest = append(dest, '"')
	dest = src.AppendFormat(dest, time.RFC3339Nano)
	dest = append(dest, '"')
	return dest
}
