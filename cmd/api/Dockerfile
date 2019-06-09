FROM golang:latest
ADD . $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog
WORKDIR $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog

RUN ls
RUN go build -o ./cmd/api/api ./cmd/api
RUN ls ./cmd/api/
EXPOSE 1234

CMD ["./cmd/api/api"]