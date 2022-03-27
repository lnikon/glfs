FROM golang:1.17

WORKDIR /workspace/glfs
RUN git clone https://github.com/lnikon/glfs .

WORKDIR /workspace/glfs-pkg
RUN git clone https://github.com/lnikon/glfs-pkg .

WORKDIR /workspace/glfs
RUN go mod tidy
RUN go build ./cmd/glfs

EXPOSE 8090

CMD ./glfs
