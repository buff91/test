FROM golang:1.15-buster as builder

WORKDIR /go/src/app
COPY . .

RUN go build -o main main.go

FROM debian:buster-slim
WORKDIR /root
RUN apt-get update && apt-get install ca-certificates -y && apt-get clean
COPY configs/config.json configs/config.json
COPY --from=builder /go/src/app/main .
# run the binary
CMD ["./main"]