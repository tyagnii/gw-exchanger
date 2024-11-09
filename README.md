# About
Exchanger is a gRPC based service provides currency rates
proto configuration repo https://github.com/tyagnii/gw-proto

## Supported currencies
RUB
EUR
USD

# Installation
```go
go get "github.com/tyagnii/gw-exchanger.git"
```

# Server
## build
Use makefile to build
`make build`

**! REMEMBER**

    The build is optimized only for Linux systems. Use your own build parameter for different OS

## run
`server serve`

## config
Configuration is placed in config.env file

# Docker-Compose Deployment
To deploy entire project with exchanger and database there is a docker-compose config file
https://github.com/tyagnii/gw-cicd