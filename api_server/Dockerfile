FROM golang:1.24

WORKDIR /usr/src/app
COPY . .


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download


RUN go build -v -o /usr/local/bin/app/cmd/api ./...


CMD ["go test ./..","app/"]
