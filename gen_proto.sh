#!/bin/bash

set -e

if [ "$BUILD_DIR" == "backend" ] || [ "$BUILD_DIR" == "frontend" ]; then

    echo "PROTO: generating..."

    # cleanup
    rm -rf backend/pb frontend/src/proto

    # create required directories
    mkdir $PWD/backend/pb
    mkdir $PWD/frontend/src/proto

    set -x

    if [ "$BUILD_DIR" == "backend" ]; then
        # generate go backend protos
        protoc heurist.proto --go_out=plugins=grpc:backend/pb
    fi

    if [ "$BUILD_DIR" == "frontend" ]; then
        # generate grpc-web client stubs
        protoc heurist.proto --js_out=import_style=commonjs,binary:frontend/src/proto  --grpc-web_out=import_style=typescript,mode=grpcweb:frontend/src/proto
    fi

    { set +x; } 2> /dev/null
    echo "PROTO: generated"

fi
