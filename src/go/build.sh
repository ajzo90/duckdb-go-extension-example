#!/bin/bash
CURRENT_DIR=$(PWD)
SRC_DIR="$1"
DUCKDB_LIB_PATH="$2"
(cd "$SRC_DIR" && go mod tidy && CGO_CFLAGS="-I${DUCKDB_LIB_PATH}/" CGO_LDFLAGS="-L${DUCKDB_LIB_PATH}" CGO_ENABLED=1 GOWORK=off go build -buildmode=c-archive -o "$CURRENT_DIR" .)