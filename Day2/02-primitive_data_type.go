package main

import "fmt"

func main() {
	var i int = 100
	var i8 int8 = 127
	var i16 int16 = 32000
	var i32 int32 = 100000
	var i64 int64 = 999999999

	// Unsigned integers
	var u uint = 200
	var u8 uint8 = 255
	var u16 uint16 = 50000
	var u32 uint32 = 100000
	var u64 uint64 = 999999999

	var f32 float32 = 10.5
	var f64 float64 = 99.99

	var b bool = true
	var r rune = 'A'
	var by byte = 'G'
	var s string = "Hello Go"

	// For all kinds of data we can use %v
	//
	// %d === integer
	// %f === float
	// %c === character
	// %t === bool
	// %s === string

	fmt.Printf("int: %v\n", i)
	fmt.Printf("int8: %v\n", i8)
	fmt.Printf("int16: %v\n", i16)
	fmt.Printf("int32: %v\n", i32)
	fmt.Printf("int64: %v\n", i64)

	fmt.Printf("uint: %v\n", u)
	fmt.Printf("uint8: %v\n", u8)
	fmt.Printf("uint16: %v\n", u16)
	fmt.Printf("uint32: %v\n", u32)
	fmt.Printf("uint64: %v\n", u64)

	fmt.Printf("float32: %v\n", f32)
	fmt.Printf("float64: %v\n", f64)

	fmt.Printf("bool: %v\n", b)
	fmt.Printf("rune: %c\n", r)
	fmt.Printf("byte: %c\n", by)
	fmt.Printf("string: %s\n", s)
}
