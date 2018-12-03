# golang_modular_app

Simple example how to build modular golang application.

## Plugin description

Example plugin implements 2 encryption algorithms : Ceasar and Verman.

It exports 3 symbols:

1. EncryptCeasar(shift int, text string  - function to encrypt text using Ceasar cipher
2. DecryptCeasar(shift int, text string  - function to encrypt text using Ceasar cipher
3. VermanCipher - variable of type vermanCipher that implements 2 methods:

- Encrypt(string) string
- Decrypt() (*string, error)

## Compile

To compile plugin run:

```
go build -buildmode=plugin -o plugin/cipher.so plugin/cipher.go
```


To compile exmaple app run:

```
go build app.go
```

To run app:

```
./app
```