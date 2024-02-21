.PHONY: build clean deploy gomodgen
all: build

_build: gomodgen
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/lambda/output/lambda lambda/main.go
	curl -LO https://github.com/SebastiaanKlippert/go-wkhtmltopdf-lambda/releases/download/0.1/lambda.zip  && unzip lambda.zip 'wkhtmltopdf' -d ../../bin/lambda/output && rm -f lambda.zip

_zip_build:
	cd ../.. && zip -r lambda.zip bin

deploy: clean build
	sls deploy --verbose

build: _build
build: _zip_build