#!/bin/bash
# no arguments, no environment variables, use config file
# expected result:
# Host:  host_from_file
# Port:  8081
# Debug: true
go run config_loader.go