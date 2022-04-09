package iterator

func Repeat(str string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += str
	}
	return result
}
