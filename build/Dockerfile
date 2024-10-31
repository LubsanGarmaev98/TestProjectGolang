FROM golang:1.23.0-alpine
# Force the go compiler to use modules
ENV GO111MODULE=on

# Create the user and group files to run unprivileged
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk update
RUN mkdir /app
WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o /app/main ./cmd/TestProjectGolang

EXPOSE 8080

USER nobody:nobody
CMD ["/app/main"]
#CMD ["ping", "www.google.com"]
