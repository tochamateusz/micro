FROM golang:1.20-alpine as base

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN 
RUN go build -v -o /usr/local/bin/app ./cmd/

FROM busybox
WORKDIR /usr/bin

COPY --from=base /usr/local/bin/app .

ENTRYPOINT [ "app" ]
