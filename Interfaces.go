package main

import (
	"bytes"
	"fmt"
)
/**
Interfaces:
	Basics, Composing interfaces, Type Conversion(The empty interface, Type switches),
	Implementing with values vs. pointers, Best practices
 */
 /**
 Best Practices:
 	Use many, small interfaces.
 		Single method interfaces are some of the most powerful and flexible:
 			io.Writer, io.Reader, interface{}
 	Don't export interfaces for types that will be consumed
 	Do export interfaces for types that will be used by package
 	Design functions and methods to receive interfaces whenever possible

  */
func main() {
	// Example 1
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	// Example 2
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// Example 3: Composing Interfaces
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello guys, this is a test!"))
	wc.Close()

	// Example 4: Type switches
	var ty interface{} = "str"
	switch ty.(type) {
	case int:
		fmt.Println("This is an int.")
	case string:
		fmt.Println("This is a string.")
	default:
		fmt.Println("I don't know the type.")
	}
}

// Example 1
type Writer interface {  // Define the interface
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Example 2
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Example 3
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
