# docker-go-example

## What this is

This is an unpolished experiment in drafting the skeleton for a containerized service in Go. It's a personal project to learn Go, so _expect_ it to be incomplete, non idiomatic, not polished, buggy or not working at all. 

## What is done

* Provides minimal image using Docker multistage builds and using single statically linked binary
* Provides, if desired, the code to deal with SSL and certificates
* Package manager files (go.mod and go.sum) are copied before the rest to generate a cached layer in Docker build, speeding up build times when dependencies won't change
* Creates a low privilege user during Docker build to avoid using root 

## What can be done

* Check the code at every step trying to adhere to what is considered "idiomatic Go"
* Find a proper logging library and implement proper logging. It must be structured (i.e. have many logging levels) and easily configurable for the minimum log level to use per environment, and provide easily CLI and JSON formatters.
* Learn how to structure properly middleware for routes
* Settle for a http client/router
* Implement failsafes if we communicate with external services
* Current build approach for the go binary makes no sense; either one disables CGO with the environment variable, _or_ uses the many flags after `go build`
