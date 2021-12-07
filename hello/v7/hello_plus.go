package main

func HelloPlus(name, address string) string {

	if name == "" {
		name = "Shane Lea"
	}
	return getPrefix(address) + name
}

func getPrefix(address string) (prefix string) {
	switch address {
	case "shanghai":
		prefix = shanghaiAddress
	case "shenzhen":
		prefix = shenzhenAddress
	default:
		prefix = beijingAddress
	}
	return
}
