default: build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o tflint-ruleset-onename

.PHONY: install
install: build
	mkdir -p ./.tflint.d/plugins
	mv -f ./tflint-ruleset-onename ./.tflint.d/plugins
