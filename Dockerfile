FROM golang:1.18-alpine

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build main.go

EXPOSE 8080

CMD [ "./main" ]