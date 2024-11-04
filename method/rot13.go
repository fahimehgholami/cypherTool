package method

import "fmt"

func EncryptROT13(input string) string {
	var encryptedMessage string
	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			encryptedMessage += string((ch-'a'+13)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			encryptedMessage += string((ch-'A'+13)%26 + 'A')
		} else if (ch >= ' ' && ch <= '@') || (ch >= '[' && ch <= '`') || (ch >= '{' && ch <= '~') {
			encryptedMessage += string(ch)
		} else {
			fmt.Printf("Error: Invalid input character detected: %v\n", string(ch))
			return ""
		}
	}
	return encryptedMessage
}
