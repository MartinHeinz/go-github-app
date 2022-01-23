#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

GO_BIN=$(command -v go)
if [[ -z "${GO_BIN}" ]]; then
    echo "go binary not found"
    exit 1
fi

# Check if OS, architecture and application version variables are set in Makefile
if [ -z "${OS:-}" ]; then
    echo "OS must be set"
    exit 1
fi
if [ -z "${ARCH:-}" ]; then
    echo "ARCH must be set"
    exit 1
fi
if [ -z "${VERSION:-}" ]; then
    echo "VERSION must be set"
    exit 1
fi

# fetch valid GOOS and GOARCH values
GOOS_OPTIONS=$("${GO_BIN}" tool dist list | cut -d '/' -f 1 | uniq | xargs)
GOARCH_OPTIONS=$("${GO_BIN}" tool dist list | cut -d '/' -f 2 | uniq | xargs)

# check if the OS value is valid GOOS value
if [[ ! " ${GOOS_OPTIONS[*]} " =~ " ${OS} " ]]; then
    echo "\"${OS}\" is not a valid OS option. valid options are: ${GOOS_OPTIONS}"
    exit 1
fi

# check if the ARCH value is valid GOARCH value
if [[ ! " ${GOARCH_OPTIONS[*]} " =~ " ${ARCH} " ]]; then
    echo "\"${ARCH}\" is not a valid ARCH option. valid options are: ${GOARCH_OPTIONS}"
    exit 1
fi

# Disable C code, enable Go modules
export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"
export GO111MODULE=on
export GOFLAGS="-mod=vendor"

# Build the application. `-ldflags -X` sets version variable in importpath for each `go tool link` invocation
"${GO_BIN}" install                                                      \
    -installsuffix "static"                                     \
    -ldflags "-X $(go list -m)/pkg/version.VERSION=${VERSION}"  \
    ./...
