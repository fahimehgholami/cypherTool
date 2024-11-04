# CypherTool
## Overview

CypherTool is a command-line utility for encrypting and decrypting text using various cipher methods. The program provides a menu-driven interface, allowing users to select from options like ROT13, Reverse Alphabet, and Vigenère ciphers. CypherTool also handles basic input validation and error handling.

Note: All numbers and punctuation symbols remain unchanged during encryption and decryption across all ciphers.

## Features

* ROT13 Encryption/Decryption: Shifts each letter by 13 positions in the alphabet.
* Reverse Alphabet Encryption/Decryption: Maps each letter to its reverse counterpart in the alphabet.
* Vigenère Encryption/Decryption: Uses a keyword to perform character-based shifts, allowing for a more complex encryption scheme.
* Interactive Menu: A user-friendly menu for selecting encryption or decryption options and navigating the tool.
* Input Validation: Prevents invalid characters and displays error messages for unsupported inputs.

## Directory Structure

* menu: Contains the Menu function that displays the main interface and handles user input.
* method: Holds functions for each encryption/decryption method.
* main.go: The main entry point that initializes the program.


## Usage

* Launch CypherTool: Run go run main.go.
* Select Option:
    1. Decrypt text
    2. Encrypt text
    3. Quit the program
* Choose Cipher Method:
    1. ROT13
    2. Reverse Alphabet
    3. Vigenère
    4. Go back to the main menu
* Enter Text: Enter the text you wish to encrypt or decrypt.
* (For Vigenère) Enter a keyword for encryption or decryption.

## Example

### Encrypting with ROT13:
    Select 2 for Encrypt.
    Choose 1 for ROT13.
    Enter your text, e.g., "Hello".
    The tool will output: Uryyb.

### Decrypting with Vigenère:
    Select 1 for Decrypt.
    Choose 3 for Vigenère.
    Enter your text, e.g., "Lxfopvefrnhr".
    Enter your keyword, e.g., "LEMON".
    The tool will output the decrypted text.

## Code Structure
#### menu/Menu.go

* Menu: Displays the main menu and handles primary navigation.
* askForMethod: Shows the available encryption methods.
* encrypt and decrypt: Handles encryption and decryption based on the user’s chosen method.
* getInput: Processes user input and removes unnecessary spaces.
* errorHandler: Manages any input errors that may occur.

#### method

* EncryptROT13: Encrypts using the ROT13 cipher.
* EncryptReverse: Encrypts by reversing the alphabet.
* EncryptVigenere: Encrypts using the Vigenère cipher with a specified keyword.
* DecryptVigenere: Decrypts Vigenère-encrypted text.
* keyToInt: Helper function to convert a key character into an integer for Vigenère calculations.