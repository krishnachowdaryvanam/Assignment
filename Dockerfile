FROM golang:latest

WORKDIR /app

COPY . .

# Build the Go application
RUN go build -o medical_vital_management

# Run the compiled binary
CMD ["./medical_vital_management"]
