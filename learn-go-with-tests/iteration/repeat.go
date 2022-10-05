package iteration

func Repeat(character string, cnt int) string {
	var repeated string
	for i := 0; i < cnt; i++ {
		repeated += character
	}
	return repeated
}
