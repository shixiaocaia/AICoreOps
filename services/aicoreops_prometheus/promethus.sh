#!/bin/bash
protoc prometheus_rpc.proto --go_out=. --go-grpc_out=.