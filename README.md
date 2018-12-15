# golang_modular_app

Simple example how to build modular golang application.

## Plugin description

Example plugin implements 2 encryption algorithms : Caesar and Verman.

It exports 3 symbols:

- EncryptCaesar(int, string)  - function to encrypt text using Caesar cipher
- DecryptCaesar(int, string)  - function to decrypt text using Caesar cipher
- VermanCipher - variable of type vermanCipher that implements 2 methods:
    - Encrypt(string) string
    - Decrypt() (*string, error)

## Compile & Run

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