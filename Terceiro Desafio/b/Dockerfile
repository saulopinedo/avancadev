FROM golang:1.15.4

WORKDIR /go/src/microservice-b

COPY . .

RUN GOOS=linux go get "github.com/joho/godotenv"
RUN GOOS=linux go get "github.com/satori/go.uuid"
RUN GOOS=linux go get "github.com/streadway/amqp"
RUN GOOS=linux go get "github.com/wesleywillians/go-rabbitmq/queue"

RUN GOOS=linux go build

EXPOSE 9091

ENTRYPOINT [ "./microservice-b" ]