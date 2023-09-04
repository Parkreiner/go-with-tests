package main

import "fmt"

const defaultLocaleKey = "en"
const defaultRecipient = "world"

var greetings = map[string][2]string{
	"en": {"Hello, ", "!"},
	"es": {"Hola, ", "!"},
	"fr": {"Bonjour, ", "!"},
}

func getGreetingInfo(locale string) (prefix string, suffix string) {
	greetingInfo, ok := greetings[locale]
	if !ok {
		greetingInfo = greetings[defaultLocaleKey]
	}

	return greetingInfo[0], greetingInfo[1]
}

func Hello(name string, locale string) string {
	prefix, suffix := getGreetingInfo(locale)

	target := name
	if target == "" {
		target = defaultRecipient
	}

	return prefix + target + suffix
}

func main() {
	msg := Hello("World", "en")
	fmt.Println(msg)
}
