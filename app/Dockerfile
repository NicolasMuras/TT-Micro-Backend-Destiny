FROM golang:alpine AS build-golang

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

RUN apk add git
# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o destiny-api ./cmd/


FROM alpine:3.16

WORKDIR /Destiny
# Add logs folder
RUN mkdir logs

# Copy Main app executable
COPY --from=build-golang /build/destiny-api ./

EXPOSE 8080

CMD ["/Destiny/destiny-api"]