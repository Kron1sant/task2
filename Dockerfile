FROM golang:1.18 as builder

WORKDIR /build

COPY go.mod ./  
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /build/bin/task2 ./...

FROM alpine:3
COPY --from=builder /build/bin/task2 /bin/task2

VOLUME [ "/data" ]

ENTRYPOINT ["/bin/task2"]