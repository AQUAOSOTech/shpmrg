deps:
	go get github.com/jonas-p/go-shp

build:
		rm -rf build
		mkdir build

		env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/shpmrg_linux
		env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o build/shpmrg_mac
		env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/shpmrg.exe

		chmod +x build/shpmrg_linux
		chmod +x build/shpmrg_mac
.PHONY: build
