FROM golang:1.18-alpine

COPY . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]