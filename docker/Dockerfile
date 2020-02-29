# ----------------------------------------------
# Build stage
# ----------------------------------------------

FROM golang:1.13-alpine AS build

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# Add unprivileged user. See https://stackoverflow.com/a/55757473/1403638 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# cache dependencies
WORKDIR /application
ADD ./go.mod .
ADD ./go.sum .
RUN go mod download

# do actual build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' main.go

# ----------------------------------------------
# Certificates stage (if needed)
# ----------------------------------------------

# FROM alpine:3.11 as certs
# RUN apk add ca-certificates && update-ca-certificates

# ----------------------------------------------
# Application stage
# ----------------------------------------------

FROM scratch

# Import the user and group files from the builder
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

# Copy certificates, if needed
# COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy main app
COPY --from=build /application/main /application/main 

# Select unprivileged user
USER appuser:appuser

ENTRYPOINT [ "/application/main" ]
