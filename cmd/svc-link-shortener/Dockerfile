FROM golang:1.16.5-alpine3.14

RUN mkdir /app
RUN mkdir /project

WORKDIR /project
COPY . .

RUN go build -o /app/main cmd/svc-link-shortener/main.go

WORKDIR /app
RUN rm -rf /project

CMD ["./main"]

