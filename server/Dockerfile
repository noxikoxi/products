FROM golang:1.24.2

RUN useradd -ms /bin/bash student

WORKDIR /home/student/

ENV CGO_ENABLED=1

COPY go.mod go.sum /home/student/
COPY controllers/ /home/student/controllers/
COPY database/ /home/student/database/
COPY models/ /home/student/models/
COPY routers/ /home/student/routers/
COPY server.go /home/student/

RUN chown -R student:student /home/student/database/ \
    && chmod -R 755 /home/student/database/ \
    && chmod 644 /home/student/database/database.db \
    && go build -o myapp .

EXPOSE 1323

USER student

CMD ["./myapp"]