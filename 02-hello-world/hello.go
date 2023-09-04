package main

import "fmt"

var greetings = map[string][2]string{
	"en": {"Hello, ", "!"},
	"es": {"Hola, ", "!"},
}

func Hello(name string, locale string) string {
	greetingInfo, ok := greetings[locale]
	if !ok {
		greetingInfo = greetings["en"]
	}

	prefix, suffix := greetingInfo[0], greetingInfo[1]

	target := name
	if target == "" {
		target = "world"
	}

	return prefix + target + suffix
}

func main() {
	msg := Hello("World", "en")
	fmt.Println(msg)
}
