default: build

test:
	go test ./...

build:
	go build -o tflint-ruleset-onename

install: build
	mkdir -p ./.tflint.d/plugins
	mv ./tflint-ruleset-onename ./.tflint.d/plugins
