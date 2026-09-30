// Task 2: All major Go primitive data types.
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("=== Signed integers ===")
	var i8 int8 = math.MaxInt8
	var i16 int16 = math.MaxInt16
	var i32 int32 = math.MaxInt32
	var i64 int64 = math.MaxInt64
	var i int = 42 // 64-bit on 64-bit systems
	row("int8", i8, unsafe.Sizeof(i8), "-128 to 127")
	row("int16", i16, unsafe.Sizeof(i16), "-32,768 to 32,767")
	row("int32", i32, unsafe.Sizeof(i32), "about ±2.1 billion")
	row("int64", i64, unsafe.Sizeof(i64), "about ±9.2 quintillion")
	row("int", i, unsafe.Sizeof(i), "same as int64 on 64-bit")

	fmt.Println("\n=== Unsigned integers (no negatives) ===")
	var u8 uint8 = math.MaxUint8
	var u16 uint16 = math.MaxUint16
	var u32 uint32 = math.MaxUint32
	var u64 uint64 = math.MaxUint64
	var u uint = 7
	row("uint8", u8, unsafe.Sizeof(u8), "0 to 255")
	row("uint16", u16, unsafe.Sizeof(u16), "0 to 65,535")
	row("uint32", u32, unsafe.Sizeof(u32), "0 to about 4.3 billion")
	row("uint64", u64, unsafe.Sizeof(u64), "0 to about 18.4 quintillion")
	row("uint", u, unsafe.Sizeof(u), "same as uint64 on 64-bit")

	fmt.Println("\n=== Floating point ===")
	var f32 float32 = 3.14159265358979
	var f64 float64 = 3.14159265358979
	row("float32", f32, unsafe.Sizeof(f32), "~7 digits of precision")
	row("float64", f64, unsafe.Sizeof(f64), "~15 digits (default for decimals)")

	fmt.Println("\n=== Complex numbers ===")
	var c64 complex64 = complex(1, 2)
	c128 := complex(3.5, -4.0)
	row("complex64", c64, unsafe.Sizeof(c64), "float32 real + imag")
	row("complex128", c128, unsafe.Sizeof(c128), "float64 real + imag")
	fmt.Printf("   real(c128)=%v imag(c128)=%v\n", real(c128), imag(c128))

	fmt.Println("\n=== Boolean, string, byte, rune ===")
	var isPermanent bool = true
	var name string = "Priya"
	var b byte = 'A' // alias for uint8: one ASCII character / raw byte
	var r rune = 'ಕ' // alias for int32: one Unicode code point
	row("bool", isPermanent, unsafe.Sizeof(isPermanent), "true or false")
	row("string", name, unsafe.Sizeof(name), "header size; text is UTF-8 bytes")
	row("byte", fmt.Sprintf("%d (%c)", b, b), unsafe.Sizeof(b), "alias of uint8")
	row("rune", fmt.Sprintf("%d (%c)", r, r), unsafe.Sizeof(r), "alias of int32")

	fmt.Println("\n=== Overflow wraps around (a common bug) ===")
	var small uint8 = 255
	small++
	fmt.Println("uint8 255 + 1 =", small)
	var tiny int8 = 127
	tiny++
	fmt.Println("int8  127 + 1 =", tiny)

	fmt.Println("\n=== Float precision ===")
	x, y := 0.1, 0.2
	fmt.Println("0.1 + 0.2 =", x+y, "(not exactly 0.3, so never compare money floats with ==)")
}

func row(name string, v any, size uintptr, note string) {
	fmt.Printf("%-11s %-22s %2d bytes  %s\n", name, fmt.Sprint(v), size, note)
}
