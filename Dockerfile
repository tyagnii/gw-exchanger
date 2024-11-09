FROM golang:1.23.2 as build

WORKDIR /usr/gw-exchanger

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

FROM scratch
WORKDIR /
COPY --from=build /usr/gw-exchanger/bin/server /
COPY --from=build /usr/gw-exchanger/config.env /

CMD ["/server", "serve"]

