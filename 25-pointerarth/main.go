package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := []any{10, 14.34, false, 'A', "Hello World", 10}
	fmt.Println(arr)
	var ptr *any = &arr[0]
	fmt.Println(*ptr)

	for i := 1; i < len(arr); i++ {
		var uptr uintptr = uintptr(unsafe.Pointer(ptr))
		uptr += unsafe.Sizeof(arr[i])
		ptr = (*any)(unsafe.Pointer(uptr))
		fmt.Println(*ptr)
	}

}

// if a and b are uint8-uint64 and int8-int64 or float32 or float64 it should return a+b -> float64 ,nil
