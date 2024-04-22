package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat(character string, no_of_times int) string {
	var repeated string
	for i :=0; i < no_of_times; i++ {
		repeated += character
	}
	return repeated
}