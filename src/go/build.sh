#!/bin/bash
CURRENT_DIR=$(pwd)
echo "Inside build.sh '$DUCKDB_PLATFORM' for the quack extension."
echo $CURRENT_DIR


go env GOARCH GOOS


if [ "$DUCKDB_PLATFORM" == "osx_amd64" ]; then
export GOOS=darwin
export GOARCH=amd64
elif [ "$DUCKDB_PLATFORM" == "osx_arm64" ]; then
export GOOS=darwin
export GOARCH=arm64
elif [ "$DUCKDB_PLATFORM" == "linux_amd64_gcc4" ]; then
export GOOS=linux
export GOARCH=amd64
elif [ "$DUCKDB_PLATFORM" == "linux_arm64" ]; then
export GOOS=linux
export GOARCH=arm64
elif [ "$DUCKDB_PLATFORM" == "linux_amd64" ]; then
export GOOS=linux
export GOARCH=amd64
else
  exit 1
fi

go env GOARCH GOOS


SRC_DIR="$1"
DUCKDB_LIB_PATH=$(realpath "$2")
ls -l $DUCKDB_LIB_PATH
(cd "$SRC_DIR" && GOOS=$GOOS GOARCH=$GOARCH go mod tidy && GOOS=$GOOS GOARCH=$GOARCH CGO_CFLAGS="-I${DUCKDB_LIB_PATH}/" CGO_LDFLAGS="-L${DUCKDB_LIB_PATH}" CGO_ENABLED=1 GOWORK=off go build -buildmode=c-archive -o "$CURRENT_DIR" .)
echo "Whats in current dir?"
ls -sl "$CURRENT_DIR"