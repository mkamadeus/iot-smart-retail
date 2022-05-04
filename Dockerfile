FROM golang:alpine

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o iot

ENTRYPOINT ["/app/iot"]

# build
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o iot .

# run
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/iot .

CMD ["./iot"]