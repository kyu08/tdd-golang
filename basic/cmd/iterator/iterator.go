package iterator

func Repeat(str string, roopNumber int) string {
	var result string
	for i := 0; i < roopNumber; i++ {
		result += str
	}
	return result
}
