package repeat

func Repeat(character string, forTime int) string {
	var repeated string
	for i := 0; i < forTime; i++ {
		repeated += character
	}
	return repeated
}
