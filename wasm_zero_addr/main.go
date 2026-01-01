package main

import "unsafe"

//export test_zero
func test_zero() int32 {
	ptr := (*int32)(unsafe.Pointer(uintptr(0)))
	*ptr = 2025
	return *ptr
}

func main() {}
