#!/bin/bash
# no arguments, use environment variables, ignore config file
# expected result:
# Host:  host_from_env
# Port:  8082
# Debug: false
APP_HOST="host_from_env" APP_PORT=8082 APP_DEBUG=false \
    go run config_loader.go