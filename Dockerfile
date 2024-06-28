# syntax=docker/dockerfile:1
FROM golang:1.21.8

# Set destination for COPY.
WORKDIR /assignment-amori

# Copy the source code.
COPY / ./

# Build.
RUN go build -o /assignment-amori-http cmd/app/http/*.go

# Run Application.
CMD [ "/bin/sh", "-c", "/assignment-amori-http"]