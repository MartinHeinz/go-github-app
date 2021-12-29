FROM golang:1.17-alpine

RUN apk add gcc g++

# https://github.com/mattn/go-sqlite3/issues/822
ENV CGO_CFLAGS="-g -O2 -Wno-return-local-addr"
# Docker image for running tests. This image is needed because tests use SQLite3 as in-memory database
# and that requires CGO to be enabled, which in turn requires GCC and G++ to be installed.
