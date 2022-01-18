#!/bin/bash
GOOS=js GOARCH=wasm go build -o noita.wasm noita-go
