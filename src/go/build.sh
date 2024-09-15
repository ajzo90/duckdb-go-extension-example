#!/bin/bash
CURRENT_DIR=$(PWD)
echo $CURRENT_DIR
SRC_DIR="$1"
DUCKDB_LIB_PATH=$(realpath "$2")
ls -l $DUCKDB_LIB_PATH
(cd "$SRC_DIR" && go mod tidy && CGO_CFLAGS="-I${DUCKDB_LIB_PATH}/" CGO_LDFLAGS="-L${DUCKDB_LIB_PATH}" CGO_ENABLED=1 GOWORK=off go build -buildmode=c-archive -o "$CURRENT_DIR" .)
echo "Whats in current dir?"
ls -sl "$CURRENT_DIR"