package reverse

func Cipher(message string, key string, decrypt bool) string {
	return reverse(message)
}

func reverse(message string) string {
	runes := []rune(message)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
