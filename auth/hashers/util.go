package hashers

import (
	"math/rand"
)

func MaskHash(hash string, show int, char string) string {
	masked := hash[:show]

	length := len(hash)
	for i := 0; i < length-show; i++ {
		masked += char
	}

	return masked
}

func ConstantTimeCompare(val1 string, val2 string) bool {
	if len(val1) != len(val2) {
		return false
	}

	length := len(val1)
	result := 0

	for i := 0; i < length; i++ {
		result |= int(val1[i]) ^ int(val2[i])
	}

	return result == 0
}

func RandomString(length int) string {
	return RandomStringWithChars(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func RandomStringWithChars(length int, chars string) string {
	var result string

	choices := []byte(chars)

	for i := 0; i < length; i++ {
		result += string(choices[rand.Intn(len(chars))])
	}

	return result
}
