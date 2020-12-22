# Start from the latest golang base image
FROM golang:latest as builder
LABEL maintainer="Joseph Bironas <josebiro@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/tod2rgb

######## Start a new stage #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

#Eventually, for metrics
EXPOSE 9090

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]