FROM golang:1.17

WORKDIR /workspace
RUN git clone https://github.com/lnikon/upcxx-operator

WORKDIR /workspace/glfs
RUN git clone https://github.com/lnikon/glfs .

WORKDIR /workspace/glfs
RUN go mod tidy
RUN go build ./cmd/glfs

ARG PORT=8080
ENV PORT $PORT

EXPOSE $PORT

CMD ./glfs
