# ----------------------------------------------
# Build stage
# ----------------------------------------------

FROM golang:1.13-alpine AS build

# cache dependencies
WORKDIR /application
ADD ./src/go.mod .
ADD ./src/go.sum .
RUN go mod download

# do actual build
COPY ./src .
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
COPY --from=build /application/main /application/main 
# COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/application/main" ]
