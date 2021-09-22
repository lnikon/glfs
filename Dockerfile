FROM golang:1.17

WORKDIR /app
COPY . .
RUN go mod tidy
# RUN go install -v ./..

CMD "client"
