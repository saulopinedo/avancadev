FROM golang:1.15.4

WORKDIR /go/src/microservice-c

COPY . .

RUN GOOS=linux go build

EXPOSE 9092

ENTRYPOINT [ "./microservice-c" ]