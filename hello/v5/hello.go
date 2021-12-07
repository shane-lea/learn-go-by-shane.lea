package main

const englishNamePrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "Shane Lea"
	}
	return englishNamePrefix + name
}
