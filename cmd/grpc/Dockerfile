FROM golang:latest
ADD . $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog
WORKDIR $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog

RUN ls
RUN go build -o ./cmd/grpc/grpc ./cmd/grpc
RUN ls ./cmd/grpc/
EXPOSE 10001

CMD ["./cmd/grpc/grpc"]