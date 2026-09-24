// Task 3: Zero values. In Go every variable declared without a value
// gets a safe default (its "zero value"), so there is never garbage memory.
package main

import "fmt"

// Employee shows that a struct's zero value is all of its fields at zero.
type Employee struct {
	ID     int
	Name   string
	Salary float64
	Active bool
}

func main() {
	var i int
	var f float64
	var s string
	var b bool
	var r rune
	var by byte
	var c complex128
	var p *int
	var sl []string
	var m map[string]int
	var arr [3]int
	var e Employee
	var fn func()
	var ch chan int
	var iface any
	var err error

	fmt.Println("Type            Zero value")
	fmt.Println("--------------  ----------")
	show("int", i)
	show("float64", f)
	fmt.Printf("%-15s %q  (empty string)\n", "string", s)
	show("bool", b)
	show("rune", r)
	show("byte", by)
	show("complex128", c)
	show("*int (pointer)", p)
	show("[]string", sl)
	show("map[string]int", m)
	show("[3]int", arr)
	fmt.Printf("%-15s %+v\n", "struct", e)
	fmt.Printf("%-15s nil = %v\n", "func()", fn == nil)
	show("chan int", ch)
	show("any", iface)
	show("error", err)

	fmt.Println("\n=== Checking for nil ===")
	fmt.Println("pointer is nil:", p == nil)
	fmt.Println("slice   is nil:", sl == nil, " len:", len(sl))
	fmt.Println("map     is nil:", m == nil, " len:", len(m))
	fmt.Println("error   is nil:", err == nil)

	fmt.Println("\n=== What you can safely do with zero values ===")
	sl = append(sl, "appending to a nil slice works")
	fmt.Println("OK:", sl[0])
	fmt.Println("OK: reading a nil map returns the zero value:", m["missing"])
	e.Name = "Arjun" // a zero struct is ready to use
	fmt.Printf("OK: a zero struct is usable: %+v\n", e)

	fmt.Println("\n=== What panics ===")
	writeNilMap()
	derefNilPointer()
	fmt.Println("Fix: create maps with make(map[string]int) and check pointers for nil.")
}

func show(t string, v any) { fmt.Printf("%-15s %v\n", t, v) }

func writeNilMap() {
	defer func() { fmt.Println("PANIC: writing to a nil map ->", recover()) }()
	var m map[string]int
	m["x"] = 1
}

func derefNilPointer() {
	defer func() { fmt.Println("PANIC: dereferencing a nil pointer ->", recover()) }()
	var p *int
	fmt.Println(*p)
}
