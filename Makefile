# Build an Exchanger service
.PHONY: build
build:
	@ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/server ./main.go