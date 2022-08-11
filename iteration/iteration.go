package iteration

func Repeat(chr string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += chr
	}
	return repeated
}
