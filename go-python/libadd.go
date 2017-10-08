package main

import "C"

//export add
func add(left, right *C.char) *C.char {
	// return left + right
	merge := C.GoString(left) + C.GoString(right)
	return C.CString(merge)
}

func main() {}
