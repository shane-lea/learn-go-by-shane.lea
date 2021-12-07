package main

import "fmt"

const beijingAddress = ":Beijing"
const shanghaiAddress = ":Shanghai"
const shenzhenAddress = "Shenzhen"

func Hello(name, address string) string {

	if name == "" {
		name = "Shane Lea"
	}

	adr := beijingAddress

	switch address {
	case "shanghai":
		adr = shanghaiAddress
	case "shenzhen":
		adr = shenzhenAddress
	}
	return name + adr
}

func main() {
	fmt.Println(Hello("", "shanghai"))
}
