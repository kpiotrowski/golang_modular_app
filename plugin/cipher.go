package main

import "errors"
import "math/rand"

//EncryptCeasar replace every letter in the `text` by a letter shifted by `shift`.
func EncryptCeasar(shift int, text string) string {
	return ceasarShift(shift, text)
}

//DecryptCeasar replace every letter in the `text` by a letter shifted by `-shift`.
func DecryptCeasar(shift int, text string) string {
	return ceasarShift(-shift, text)
}

//VermanCipher is similar to the Ceasar sipher, however it shift every letter by a random number of positions
type vermanCipher struct {
	encryptedText *string
	key           []rune
}

//VermanCipher is a exported variable that will be shared with external app
var VermanCipher vermanCipher

//Encrypt generates encryption key and use it in the ceasar shift method
func (v *vermanCipher) Encrypt(text string) string {
	v.key = generateShiftKey(len(text))
	encryptedText := vermanShift(v.key, text, false)
	v.encryptedText = &encryptedText
	return encryptedText
}

//Decrypt uses already generated key to decrypt text
func (v *vermanCipher) Decrypt() (*string, error) {
	if v.encryptedText == nil || len(v.key) == 0 {
		return nil, errors.New("you first need to encrypt value")
	}
	decryptedText := vermanShift(v.key, *v.encryptedText, true)
	return &decryptedText, nil
}

func ceasarShift(shift int, text string) string {
	runes := make([]rune, len(text))
	for i, r := range text {
		runes[i] = shiftLetter(r, shift)
	}
	return string(runes)
}

func vermanShift(shift []rune, text string, decrypt bool) string {
	factor := 1
	if decrypt {
		factor = -1
	}

	runes := make([]rune, len(text))
	for i, s := range text {
		runes[i] = shiftLetter(s, factor*int(shift[i]))
	}
	return string(runes)
}

func shiftLetter(letter rune, shift int) rune {
	letterCode := int(letter) + (shift % 26)
	if letterCode > int('~') {
		letterCode -= 95
	}
	if letterCode < int(' ') {
		letterCode += 95
	}
	return rune(letterCode)
}

func generateShiftKey(size int) []rune {
	runes := make([]rune, size)
	for i := 0; i < size; i++ {
		runes[i] = rune(rand.Int() % 95)
	}
	return runes
}
