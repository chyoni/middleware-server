FROM golang:1.19-alpine

# RUN apk add build-base

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

# COPY *.go ./

COPY . /app

RUN go build -o /middleware-server

EXPOSE 8081

CMD [ "/middleware-server" ]