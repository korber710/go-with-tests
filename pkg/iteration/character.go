package iteration

func Repeat(character string, repeatCount int) string {
	characters := ""
	if repeatCount == 0 {
		repeatCount = 1
	}
	for i := 0; i < repeatCount; i++ {
		characters += character
	}
	return characters
}
