.PHONY: build buildPlugin clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/cipherEngine cipherEngine/main.go

buildPlugin:
	env GOOS=linux go build -ldflags="-s -w" -buildmode=plugin -o bin/plugin/cipher.so ../plugin/cipher.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean buildPlugin build
	sls deploy --verbose
