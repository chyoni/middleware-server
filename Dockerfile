FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /middleware-server

EXPOSE 8081

CMD [ "/middleware-server" ]