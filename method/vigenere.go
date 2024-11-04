package method

import "fmt"

func EncryptVigenere(input, key string) string {

	var keyCh rune
	var count int
	var result string

	for _, ch := range input {
		if !(ch >= '!' && ch <= '~') {
			fmt.Printf("Error: Invalid input character detected: %v\n", string(ch))
			return ""
		}
		if count >= len(key) {
			count = 0
		}
		keyCh = rune(key[count])
		count++

		spacesToMove := keyToInt(keyCh)
		if ch >= 'a' && ch <= 'z' {
			ch += spacesToMove
			if ch > 'z' {
				ch -= 26
			}
		}
		if ch >= 'A' && ch <= 'Z' {
			ch += spacesToMove
			if ch > 'Z' {
				ch -= 26
			}
		}

		result += string(ch)
	}

	return result
}
func keyToInt(ch rune) rune {

	if ch >= 'a' && ch <= 'z' {
		return ch - 'a'
	}

	if ch >= 'A' && ch <= 'Z' {
		return ch - 'A'
	}
	return 0
}

func DecryptVigenere(input, key string) string {

	var keyCh rune
	var count int
	var result string

	for _, ch := range input {
		if !(ch >= '!' && ch <= '~') {
			fmt.Printf("Error, invalid input at %v\n", string(ch))
			return ""
		}
		if count >= len(key) {
			count = 0
		}
		keyCh = rune(key[count])
		count++

		spacesToMove := keyToInt(keyCh)
		if ch >= 'a' && ch <= 'z' {
			ch -= spacesToMove
			if ch < 'a' {
				ch += 26
			}
		}
		if ch >= 'A' && ch <= 'Z' {
			ch -= spacesToMove
			if ch < 'A' {
				ch += 26
			}
		}

		result += string(ch)
	}

	return result
}
