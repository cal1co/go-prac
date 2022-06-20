package main

import (
	"fmt"
)

func main() {
	// var n bool = true
	// fmt.Printf("%v, %T\n", n, n)

	// n := 1 == 1
	// m := 1 == 2
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T\n", m, m)

	// a := 8 // 2^3
	// fmt.Println(a << 3)
	// fmt.Println(a >> 3)

	// a := 10.2
	// b := 3.7
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// var n complex64 = 2i
	// fmt.Printf("%v, %T\n", n, n)

	// a := 1 + 2i
	// b := 2 + 5.2i
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// var n complex64 = 1 + 2i
	// fmt.Printf("%v, %T\n", real(n), real(n))
	// fmt.Printf("%v, %T\n", imag(n), imag(n))

	// var n complex128 = complex(5, 12)
	// fmt.Printf("%v, %T\n", n, n)

	// s := "this is a string"
	// fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// s := "this is a string"
	// s2 := "this is also a string"
	// fmt.Printf("%v, %T\n", s+s2, s+s2)

	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

// Primitives:

// - Boolean type:
// Values are true or false
// Not an alias for other types
// Zero value is false

// - Numeric types:

// -- Integers --
// * Signed Integers
// 	*-int type has varying size, but min 32 bits
// 	*-8 bit (int8) through 64(int64)
// * Unsigned integers
// 	*-8 bit (byte and uint8) through 32 bit (uint32)
// * Arithmetic operations
// 	*- Addition, subtraction, multiplications, division, remainder
// * Bitwise opeartions
// 	*-And, or, xor, and not
// * Zero value is 0
// * Can't mix types in same family! (uint16+uint32=error)

// -- Floating point numbers --
// * Follow IEEE-754 standard
// * Zero value is 0
// * 32 and 64 bit versions
// * Literal types
// 	*-Decimal
// 	*-Exponent (13e18 or 2E10)
// 	*-Mixed(13.7e12)
// * Arithmetic operations
//	*-Addition, subtraction, multiplication, division

// -- Complex Numbers --
// * Zero value is (0+0i)
// * 64 and 128 bit versions
// * Built-in functions
// 	*-complex - make complex number from two floats
// 	*-real - get real part as float
// 	*-imag - get imaginary part as float
// * Arithmetic operations
// 	*-Addition, subtraction, multiplication, division

// -- Text types --
// * Strings
// 	*-UTF-8
// 	*-Immutable
// 	*-Can be concatenated with plus (+) operator
// 	*-Can be converted to []byte
// * Rune
// 	*-UTF-32
// 	*-Alis for int32
// 	*-Special methods normally required to process
// 		** eg strings.Reader#ReadRune
