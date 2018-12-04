package main

import (
	"context"
	"plugin"

	"github.com/aws/aws-lambda-go/lambda"
)

//LambdaResponse contains response structure
type LambdaResponse struct {
	PlainText   string
	EncryptedC4 string
	DecryptedC4 string
	EncryptedV  string
	DecryptedV  string
}

//Define interface with the same set of methods as vermanCipher structure
type encryptionEngine interface {
	Encrypt(string) string
	Decrypt() (*string, error)
}

//Handler is function executed by lambda engine
func Handler(ctx context.Context) (LambdaResponse, error) {
	resp := LambdaResponse{}

	// Load Cipher plugin
	pluginModule, err := plugin.Open("/opt/cipher.so")
	if err != nil {
		return resp, err
	}
	//Load EncryptCeasar function
	encryptCeasarSymbol, err := pluginModule.Lookup("EncryptCeasar")
	if err != nil {
		return resp, err
	}
	//Load DecryptCeasar function
	decryptCeasarSymbol, err := pluginModule.Lookup("DecryptCeasar")
	if err != nil {
		return resp, err
	}
	//Load VermanCipher variable
	vermanCipherSymbol, err := pluginModule.Lookup("VermanCipher")
	if err != nil {
		return resp, err
	}

	//Cast encryptCeasar symbol to the correct type
	encryptCeasarFunc := encryptCeasarSymbol.(func(int, string) string)
	//Cast encryptCeasar symbol to the correct type
	decryptCeasarFunc := decryptCeasarSymbol.(func(int, string) string)
	//Cast vermanCipher symbol to the correct interface type
	vermanCipherIf := vermanCipherSymbol.(encryptionEngine)

	resp.PlainText = "This is my super secret text 007!"
	resp.EncryptedC4 = encryptCeasarFunc(4, resp.PlainText)
	resp.DecryptedC4 = decryptCeasarFunc(4, resp.EncryptedC4)
	resp.EncryptedV = vermanCipherIf.Encrypt(resp.PlainText)
	decryptedV, err := vermanCipherIf.Decrypt()
	if err != nil {
		return resp, err
	}
	resp.DecryptedV = *decryptedV
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
