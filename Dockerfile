FROM golang:1.23.4-alpine AS builder

WORKDIR /MillyPhoto

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .env .

RUN go build -o MillyPhoto .

FROM alpine:latest

WORKDIR /MillyPhoto

COPY --from=builder /MillyPhoto/MillyPhoto .
COPY --from=builder /MillyPhoto/.env .
COPY --from=builder /MillyPhoto/templates ./templates
COPY --from=builder /MillyPhoto/uploads ./uploads
COPY --from=builder /MillyPhoto/rental ./study

EXPOSE 8082

CMD ["./MillyPhoto"]