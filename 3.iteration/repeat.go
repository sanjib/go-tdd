package iteration

const repeatCount = 5

func Repeat(input string) string {
	output := ""
	for i := 0; i < repeatCount; i++ {
		output += input
	}
	return output
}
