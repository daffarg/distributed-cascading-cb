FROM golang:1.21-alpine3.18 AS builder
RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /build
COPY . .
RUN go mod download

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/main .

EXPOSE 5320

RUN apk add tzdata

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime


CMD ["./main"]