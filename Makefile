# Build an Exchanger service
.PHONY: build
build:
	@ go build -o ./bin/server ./main.go