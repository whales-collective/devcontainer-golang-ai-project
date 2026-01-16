FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY main.go .

RUN <<EOF
go mod tidy 
go build
EOF


FROM scratch
WORKDIR /app
COPY --from=builder /app/dmr-agent .
ENTRYPOINT ["./dmr-agent"]
