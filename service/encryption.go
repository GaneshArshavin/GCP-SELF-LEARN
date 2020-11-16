package service

import (
	"strings"
)

func rot47(input string) string {

	var result []string
	for i := range input[:len(input)] {
		j := int(input[i])
		if (j >= 33) && (j <= 126) {
			result = append(result, string(rune(33+((j+14)%94))))
		} else {
			result = append(result, string(input[i]))
		}

	}
	return strings.Join(result, "")
}

//Encrypt
func RotEn(strRaw string) string {
	return rot47(strRaw)
}

//Decrypt
func RotDn(strRaw string) string {
	return rot47(strRaw)
}
