FROM golang:alpine

RUN mkdir /myApp
ADD . /myApp

WORKDIR /myApp
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080