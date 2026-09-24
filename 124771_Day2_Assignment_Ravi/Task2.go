package main

import "fmt"

func main() {
	var isTrue bool = true
	var i8 int8 = -126
	var i16 int16 = -3276
	var i32 int32 = -214748364
	var i64 int64 = -922337203685477580
	var i int = 41
	var u8 uint8 = 254
	var u16 uint16 = 65534
	var u32 uint32 = 4294967294
	var u64 uint64 = 18446744073709551614
	var u uint = 100
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i
	var s string = "Hello, Ravi!"
	var b byte = 'A'
	var r rune = '世'

	fmt.Println("bool: ", isTrue)
	fmt.Println("int8: ", i8)
	fmt.Println("int16: ", i16)
	fmt.Println("int32: ", i32)
	fmt.Println("int64: ", i64)
	fmt.Println("int: ", i)
	fmt.Println("uint8: ", u8)
	fmt.Println("uint16: ", u16)
	fmt.Println("uint32: ", u32)
	fmt.Println("uint64: ", u64)
	fmt.Println("uint: ", u)
	fmt.Println("float32: ", f32)
	fmt.Println("float64: ", f64)
	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)
	fmt.Println("string: ", s)
	fmt.Println("byte:", b)
	fmt.Println("rune: ", r)
}
