FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /PokerGame

EXPORSE 8080

CMD["/PokerGame"]
