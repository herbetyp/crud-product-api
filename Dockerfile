FROM golang:1.23.4-alpine3.21

WORKDIR /go/src/app

COPY . .

EXPOSE 3000

RUN go build -o main cmd/main.go

CMD [ "./main" ]
