default: build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o tflint-ruleset-naming-convention

.PHONY: install
install: build
	mkdir -p ./.tflint.d/plugins
	mv -f ./tflint-ruleset-naming-convention ./.tflint.d/plugins
