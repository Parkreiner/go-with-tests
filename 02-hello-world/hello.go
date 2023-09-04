package main

import "fmt"

const helloPrefix = "Hello, "
const helloSuffix = "!"

func Hello(name string) string {
	var target string
	if name == "" {
		target = "world"
	} else {
		target = name
	}

	return helloPrefix + target + helloSuffix
}

func main() {
	msg := Hello("World")
	fmt.Println(msg)
}
