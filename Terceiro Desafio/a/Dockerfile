FROM golang:1.15.4

WORKDIR /go/src/microservice-a

COPY . .

RUN GOOS=linux go get "github.com/joho/godotenv"
RUN GOOS=linux go get "github.com/wesleywillians/go-rabbitmq/queue"

RUN GOOS=linux go build

EXPOSE 9090

ENTRYPOINT [ "./microservice-a" ]