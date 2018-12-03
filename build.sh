# Build cipher plugin
go build -buildmode=plugin -o plugin/cipher.so plugin/cipher.go

#Build main app
go build app.go

