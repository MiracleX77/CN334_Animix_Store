FROM golang:1.22.0-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o deliveryMigrate ./delivery/migrations/deliveryMigrate.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o paymentMigrate ./payment/migrations/paymentMigrate.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o orderMigrate ./order/migrations/orderMigrate.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o transactionMigrate ./transaction/migrations/transactionMigrate.go







######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache postgresql-client

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/paymentMigrate .
COPY --from=builder /app/deliveryMigrate .
COPY --from=builder /app/orderMigrate .
COPY --from=builder /app/transactionMigrate .

EXPOSE 5001

# Command to run the executable
CMD ["./main"]