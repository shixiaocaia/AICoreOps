#!/bin/bash
protoc aicoreops_api.proto --go_out=. --go-grpc_out=.