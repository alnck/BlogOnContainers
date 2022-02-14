FROM golang:1.17

RUN mkdir /app
WORKDIR /app

COPY . .

#WORKDIR ./

RUN go mod download

RUN go build ./api.go

EXPOSE 5000

CMD ["./api"]