package method

import "fmt"

func EncryptReverse(input string) string {

	var encryptedMessage string

	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			encryptedMessage += string('z' - ch + 'a')
			continue
		} else if ch >= 'A' && ch <= 'Z' {
			encryptedMessage += string('Z' - ch + 'A')
			continue
		}
		if (ch >= '!' && ch <= '@') || (ch >= '[' && ch <= '`') || (ch >= '{' && ch <= '~') {
			encryptedMessage += string(ch)
			continue
		}
		fmt.Printf("Error: Invalid input character detected: %v\n", string(ch))
		return ""
	}
	return encryptedMessage
}
