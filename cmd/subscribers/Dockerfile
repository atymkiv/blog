FROM golang:latest
ADD . $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog
WORKDIR $GOPATH/src/github.com/atymkiv/echo_frame_learning/blog

RUN ls
RUN go build -o ./cmd/subscribers/subscribers ./cmd/subscribers
RUN ls /home/
EXPOSE 4222

CMD ["./cmd/subscribers/subscribers"]