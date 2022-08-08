FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/jbrissier/hview/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/hview

RUN ls -l /go/bin/
ENTRYPOINT ["/go/bin/hview"]


FROM alpine:3.14

COPY --from=builder /go/bin/hview /usr/local/bin/hview

ENTRYPOINT ["/usr/local/bin/hview"]

