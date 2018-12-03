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
	//Load EncryptCeasar function
	encryptCeasarSymbol, err := pluginModule.Lookup("EncryptCeasar")
	if err != nil {
		fmt.Println("Unable to load ceasar encrypt function")
		os.Exit(1)
	}
	//Load DecryptCeasar function
	decryptCeasarSymbol, err := pluginModule.Lookup("DecryptCeasar")
	if err != nil {
		fmt.Println("Unable to load ceasar decrypt function")
		os.Exit(1)
	}
	//Load VermanCipher variable
	vermanCipherSymbol, err := pluginModule.Lookup("VermanCipher")
	if err != nil {
		fmt.Println("Unable to load VermanCipher variable")
		os.Exit(1)
	}

	//Cast encryptCeasar symbol to the correct type
	encryptCeasarFunc := encryptCeasarSymbol.(func(int, string) string)
	//Cast encryptCeasar symbol to the correct type
	decryptCeasarFunc := decryptCeasarSymbol.(func(int, string) string)
	//Cast vermanCipher symbol to the correct interface type
	vermanCipherIf := vermanCipherSymbol.(encryptionEngine)

	plainText := "This is my super secret text 007!"
	fmt.Printf("Plain text: \t\t%s\n", plainText)

	encryptedC4 := encryptCeasarFunc(4, plainText)
	fmt.Printf("Encrypted C4: \t\t%s\n", encryptedC4)

	decryptedC4 := decryptCeasarFunc(4, encryptedC4)
	fmt.Printf("Decrypted C4: \t\t%s\n", decryptedC4)

	encryptedV := vermanCipherIf.Encrypt(plainText)
	fmt.Printf("Encrypted V: \t\t%s\n", encryptedV)

	decryptedV, _ := vermanCipherIf.Decrypt()
	fmt.Printf("Decrypted V: \t\t%s\n", *decryptedV)
}
