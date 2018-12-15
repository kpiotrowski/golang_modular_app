package main

import (
	"fmt"
	"os"
	"plugin"
)

//Define interface with the same set of methods as vermanCipher structure
type encryptionEngine interface {
	Encrypt(string) string
	Decrypt() (*string, error)
}

func main() {
	// Load Cipher plugin
	pluginModule, err := plugin.Open("plugin/cipher.so")
	if err != nil {
		fmt.Println("Unable to load cipher module")
		os.Exit(1)
	}
	//Load EncryptCaesar function
	encryptCaesarSymbol, err := pluginModule.Lookup("EncryptCaesar")
	if err != nil {
		fmt.Println("Unable to load caesar encrypt function")
		os.Exit(1)
	}
	//Load DecryptCaesar function
	decryptCaesarSymbol, err := pluginModule.Lookup("DecryptCaesar")
	if err != nil {
		fmt.Println("Unable to load caesar decrypt function")
		os.Exit(1)
	}
	//Load VermanCipher variable
	vermanCipherSymbol, err := pluginModule.Lookup("VermanCipher")
	if err != nil {
		fmt.Println("Unable to load VermanCipher variable")
		os.Exit(1)
	}

	//Cast encryptCaesar symbol to the correct type
	encryptCaesarFunc := encryptCaesarSymbol.(func(int, string) string)
	//Cast encryptCaesar symbol to the correct type
	decryptCaesarFunc := decryptCaesarSymbol.(func(int, string) string)
	//Cast vermanCipher symbol to the correct interface type
	vermanCipherIf := vermanCipherSymbol.(encryptionEngine)

	plainText := "This is my super secret text 007!"
	fmt.Printf("Plain text: \t\t%s\n", plainText)

	encryptedC4 := encryptCaesarFunc(4, plainText)
	fmt.Printf("Encrypted C4: \t\t%s\n", encryptedC4)

	decryptedC4 := decryptCaesarFunc(4, encryptedC4)
	fmt.Printf("Decrypted C4: \t\t%s\n", decryptedC4)

	encryptedV := vermanCipherIf.Encrypt(plainText)
	fmt.Printf("Encrypted V: \t\t%s\n", encryptedV)

	decryptedV, _ := vermanCipherIf.Decrypt()
	fmt.Printf("Decrypted V: \t\t%s\n", *decryptedV)
}
