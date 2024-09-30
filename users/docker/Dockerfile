FROM golang:1.23

WORKDIR /usr/src/users

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY users/go.mod users/go.sum ./
RUN go mod download && go mod verify

COPY users/ .
RUN go build -v -o /usr/local/bin/users users.go

CMD ["users"]
