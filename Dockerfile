FROM golang:1.17

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build .
RUN go install

EXPOSE 8090

CMD glfs
