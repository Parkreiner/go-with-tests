package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	msg := Hello("World")
	fmt.Println(msg)
}
