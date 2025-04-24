FROM golang:1.24.2-alpine

WORKDIR /app

ENV CGO_ENABLED=1

COPY . .

RUN apk add build-base \
    && go mod tidy \
    && go build -o myapp .

EXPOSE 1323

CMD ["./myapp"]