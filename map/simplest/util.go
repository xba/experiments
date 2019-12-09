package simplest

import "unsafe"

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func Hash(key uint64) uint64 {
	return uint64(memhash(unsafe.Pointer(&key), 0, 4))
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

func cas(ptr *uint64, old, new uint64) uint64
