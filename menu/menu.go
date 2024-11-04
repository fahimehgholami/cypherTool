package menu

import (
	"cypherTool/method"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() {
	fmt.Println("Hi, welcome to the Cypher Tool!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("1. Decrypt\n2. Encrypt\n3. Quit\n")
		input, err := reader.ReadString('\n')
		errorHandler(err)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			askForMethod()
			i, err := reader.ReadString('\n')
			errorHandler(err)
			i = strings.TrimSpace(i)
			decrypt(i, reader)
		case "2":
			askForMethod()
			i, err := reader.ReadString('\n')
			errorHandler(err)
			i = strings.TrimSpace(i)
			encrypt(i, reader)

		case "3":
			fmt.Println("Exiting the program...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func askForMethod() {
	fmt.Println("What encryption method do you prefer?")
	fmt.Print("1. ROT13\n2. Reverse Alphabet\n3. Vigenère\n4. Go Back\n")
}
func decrypt(methodChoice string, reader *bufio.Reader) {

	switch methodChoice {
	case "1":
		fmt.Println("Enter the text to decrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		result := method.EncryptROT13(input)
		if result != "" {
			fmt.Printf("Decrypted text (ROT13): %s\n", result)
		}
	case "2":
		fmt.Println("Enter the text to decrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		result := method.EncryptReverse(input)
		if result != "" {
			fmt.Printf("Decrypted text (Reverse Alphabet): %s\n", result)
		}
	case "3":
		fmt.Println("Enter the text to encrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		fmt.Println("Enter the Key: ")
		key := getInput(reader)
		result := method.DecryptVigenere(input, key)
		if result != "" {
			fmt.Printf("Encrypted text (Vigenère): %s\n", result)
		}
	case "4":
		//Go back
	default:
		fmt.Println("Invalid method choice.")
	}
}

func encrypt(methodChoice string, reader *bufio.Reader) {

	switch methodChoice {

	case "1":
		fmt.Println("Enter the text to encrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		result := method.EncryptROT13(input)
		if result != "" {
			fmt.Printf("Encrypted text (ROT13): %s\n", result)
		}
	case "2":
		fmt.Println("Enter the text to encrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		result := method.EncryptReverse(input)
		if result != "" {
			fmt.Printf("Encrypted text (Reverse Alphabet): %s\n", result)
		}
	case "3":
		fmt.Println("Enter the text to encrypt: ")
		input := getInput(reader)
		input = strings.TrimSpace(input)
		fmt.Println("Enter the Key: ")
		key := getInput(reader)
		result := method.EncryptVigenere(input, key)
		if result != "" {
			fmt.Printf("Encrypted text (Vigenère): %s\n", result)
		}
	case "4":
		//Go back
	default:
		fmt.Println("Invalid method choice.")
	}
}
func getInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	errorHandler(err)
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	return input
}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
