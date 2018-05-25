From golang:latest
MAINTAINER DTherHtun <dther.htun@kbzbank.com>
ENV GOPATH /go
RUN go get -u github.com/gorilla/mux && mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o hw .
CMD ["./hw"]
