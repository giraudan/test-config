#!/bin/bash
# use arguments, ignore environment variables, ignore config file
# expected result:
# Host:  host_from_cli
# Port:  8083
# Debug: true
APP_HOST="host_from_env" APP_PORT=8082 APP_DEBUG=false \
    go run config_loader.go \
    --host="host_from_cli" \
    --port=8083 \
    --debug=true
