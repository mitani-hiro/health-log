FROM golang:1.24-alpine

WORKDIR /app

RUN mkdir common
COPY common/go.mod /app/common/.
COPY common /app/common

RUN mkdir -p kafka/producer
COPY kafka/producer/go.mod /app/kafka/producer/.
COPY kafka/producer /app/kafka/producer

WORKDIR /app/kafka/producer

RUN go mod download

RUN go build -o main .

CMD [ "./main" ]
