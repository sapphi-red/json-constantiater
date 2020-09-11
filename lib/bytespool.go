package lib

import (
	"sync"
	"unsafe"
)

const defaultCap = 1024

var p = sync.Pool{
	New: func() interface{} {
		tmp := make([]byte, 0, defaultCap)
		return &tmp
	},
}

// https://golang.org/src/reflect/value.go#L193
type emptyInterface struct {
	typ *struct{}
	ptr unsafe.Pointer
}

//go:nosplit
func GetFromPool() *[]byte {
	tmp := p.Get()
	return (*[]byte)(((*emptyInterface)(unsafe.Pointer(&tmp))).ptr)
}

//go:nosplit
func PutToPool(tmpPtr *[]byte) {
	*tmpPtr = (*tmpPtr)[0:0]
	p.Put(tmpPtr)
}
