FROM golang:1.17-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o dpgraham .
CMD ["/app/dpgraham"]
