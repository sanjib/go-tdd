package iteration

func Repeat(input string, count int) string {
	output := ""
	for i := 0; i < count; i++ {
		output += input
	}
	return output
}
