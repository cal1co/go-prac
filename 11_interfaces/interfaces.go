package main

import "fmt"

// func main() {
// 	// fmt.Println("Hello, World!")
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello Go!"))
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// !! COME BACK TO THIS !!

// Interfaces:

// -- Best Practices --
// * Use many, small interfaces
// 	*- Single method interfaces are some of the most powerful and flexible
// 		** io.Writer, io.Reader, interface{}
// * Don't export interfaces for types that will be consumed
// * Do export interfaces for types that will be used by package
// * Design functions and methods to receive interfaces whenever possible

// Basics
// type Writer interface{
// 		Write([]byte)(int, error)
// }

// type ConsoleWriter struct {}

// func (cw ConsoleWriter) Writer(data[]byte)(int, error){
// 		n, err := fmt.Println(string(data))
// 		return n, err
// }

// Composing interfaces
// type Writer interface{
// 		Writer([]byte)(int, error)
// }
// type Closer interface{
// 		Close() error
// }
// type WriterCloser interface{
// 		Writer
// 		Closer
// }

// Type conversion
// var wc WriterCloser = NewBufferedWriterCloser()
// bwc := wc.(*BufferedWriterCloser)

// The empty interface and type switches
// var i interface{} = 0
// switch i.(type){
// 	case int:
// 		fmt.Println("i is an integer")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i don't know what i is")
// }

// Implementing with values vs pointers
// * Methods set of value is all methods with value receivers
// * Method set of pointer is all methods, regardless of receiver type
