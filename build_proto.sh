#!/usr/bin/env bash

proto_dir=flextime-api/v1
protoc $proto_dir/*.proto \
    --go_out=plugins=grpc:rpc/flextime_v1 \
    --proto_path=$proto_dir
