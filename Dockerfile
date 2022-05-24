FROM golang:1.17-alpine AS build
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go build -o dpgraham .

FROM alpine:latest AS production
COPY --from=build /src .
EXPOSE 8080
CMD ["./dpgraham"]
