# FROM golang:1.23 AS builder

# WORKDIR /app
# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
#     GOOS=linux go build -o go-crud .

# FROM golang:1.19-alpine

# COPY --from=builder .env /app/.env

# RUN adduser -D gouser

# COPY --from=builder /app/go-crud /app/go-crud

# RUN chown -R gouser:gouser /app
# RUN chmod +x /app/go-crud

# EXPOSE 8080

# USER gouser

# CMD ["sleep","900000000000000"]
# # CMD ["/app/go-crud"]


# Stage 1: Build
FROM golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the binary
RUN go mod tidy
RUN go build -o go-crud .

# Stage 2: Run
FROM alpine:latest

# Set Environment Variables
ENV MONGODB_USER_DB=crudInit
ENV MONGODB_URL=mongodb://mongodb:27017
ENV MONGODB_USER_COLLECTION=test_collection
ENV JWT_SECRET_KEY=blablablablablablabla

# Set the working directory
WORKDIR /app/

# Copy the binary from the build stage
COPY --from=build /app/go-crud .
COPY /.env .

# Expose the port
EXPOSE 8000

# Run the binary
CMD ["./go-crud"]