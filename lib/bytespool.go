package lib

import "sync"

const defaultCap = 1024

var p = sync.Pool{
	New: func() interface{} {
		tmp := make([]byte, 0, defaultCap)
		return &tmp
	},
}

//go:nosplit
func GetFromPool() *[]byte {
	return p.Get().(*[]byte)
}

//go:nosplit
func PutToPool(tmpPtr *[]byte) {
	*tmpPtr = (*tmpPtr)[0:0]
	p.Put(tmpPtr)
}
