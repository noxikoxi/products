FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=1

COPY controllers/ controllers/
COPY database/ database/
COPY models/ models/
COPY routers/ routers/
COPY server.go .

RUN go build -o myapp server.go

FROM debian:bookworm-slim

RUN useradd -ms /bin/bash student

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/database/database.db ./database/database.db

RUN chown -R student:student . \
    && chmod 755 myapp \
    && chmod 644 ./database/database.db

# default
ENV FRONTEND_URL=http://localhost:5137
ENV PORT=1323

EXPOSE 1323

USER student

CMD ["./myapp"]