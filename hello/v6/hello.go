package main

const englishNamePrefix = "Hello, "
const peopleName = "Shane Lea"
const beijing = ":Beijing"

func Hello(name, address string) string {

	if name == "" {
		name = "Shane Lea"
	}

	if address == "beijing" {
		return peopleName + beijing
	}
	return englishNamePrefix + name

}
