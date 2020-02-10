default: build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o tflint-ruleset-naming_convention

.PHONY: install
install: build
	mkdir -p ./.tflint.d/plugins
	mv -f ./tflint-ruleset-naming_convention ./.tflint.d/plugins


.PHONY: test-install
test-install: build
	mkdir -p example/.tflint.d/plugins
	mv -f ./tflint-ruleset-naming_convention example/.tflint.d/plugins
