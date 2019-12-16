package cache

import "unsafe"

func nearest2Power(x uint64) uint64 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	x++
	return x
}

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func Hash(key []byte) uint64 {
	data := (*stringStruct)(unsafe.Pointer(&key))
	return uint64(memhash(data.str, 0, uintptr(data.len)))
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr
