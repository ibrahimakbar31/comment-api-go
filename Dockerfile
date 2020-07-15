FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOCUSTOMENV=development \
    GOCONFIGFILENAME=config-example

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
# RUN cp /build/main .
COPY /build/main .

# Export necessary port
EXPOSE 8000

# Command to run when starting the container
CMD ["/app/main"]