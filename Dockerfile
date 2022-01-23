FROM golang:1.17-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o dpgraham .
CMD ["/app/dpgraham"]

# COPY dpgraham.go go.* /src/
# RUN CGO_ENABLED=0 go build -o /bin/hello
# FROM scratch
# COPY --from=build /bin/dpgraham /bin/dpgraham
# ENTRYPOINT ["/bin/dpgraham"]
