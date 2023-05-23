FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/cespare/reflex@latest

COPY . .

CMD ["reflex", "-r", "\\.go$", "-s", "--", "sh", "-c", "go build -o app && ./app"]