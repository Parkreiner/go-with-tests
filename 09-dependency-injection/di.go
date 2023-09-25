package main

import (
	"bytes"
	"fmt"
	"io"
)

type myValue struct {
	Width  int
	Height int
}

// It greets you!
func Greet(w io.Writer, name string) {
	r := myValue{Height: 12, Width: 14}
	fmt.Printf("%d and %d", r.Width, r.Height)

	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	temp := bytes.Buffer{}

	Greet(&temp, "Yo")
}
