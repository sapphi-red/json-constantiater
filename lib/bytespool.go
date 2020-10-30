package lib

import "sync"

const defaultCap = 1024

type linker struct {
	value *[]byte
	next  *linker
}

type bytesPool struct {
	New func() *[]byte

	shared []*[]byte
	mutex  sync.Mutex
}

func (p *bytesPool) Get() *[]byte {
	p.mutex.Lock()
	if len(p.shared) <= 0 {
		p.mutex.Unlock()
		return p.New()
	}

	lastIndex := len(p.shared) - 1
	got := p.shared[lastIndex]
	p.shared = p.shared[:lastIndex]

	p.mutex.Unlock()
	return got
}

func (p *bytesPool) Put(s *[]byte) {
	p.mutex.Lock()
	p.shared = append(p.shared, s)
	p.mutex.Unlock()
}

var p = bytesPool{
	New: func() *[]byte {
		tmp := make([]byte, 0, defaultCap)
		return &tmp
	},
}

//go:nosplit
func GetFromPool() *[]byte {
	return p.Get()
}

//go:nosplit
func PutToPool(tmpPtr *[]byte) {
	*tmpPtr = (*tmpPtr)[0:0]
	p.Put(tmpPtr)
}
