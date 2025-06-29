
#BUILD IMAGE
FROM golang:1.24 AS go_build_image_client



WORKDIR /usr/src/app
COPY . .



RUN go mod download

RUN go build -v -o /execute_go ./cmd/cli/main.go




# DEPLOY

# FROM golang:1.24

# WORKDIR /app/bin

# COPY --from=go_build_image /execute_go /app/bin/execute_go


