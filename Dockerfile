FROM golang:1.18-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .

EXPOSE 3000


## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]