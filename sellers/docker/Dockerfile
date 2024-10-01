FROM golang:1.23

WORKDIR /usr/src/sellers

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY sellers/go.mod sellers/go.sum ./
RUN go mod download && go mod verify

COPY sellers/ .
RUN go build -v -o /usr/local/bin/sellers sellers.go

CMD ["sellers"]
