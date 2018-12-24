#!/bin/bash

set -ex

# download & install protoc & protoc-plugins
if [ "$BUILD_DIR" == "backend" ] || [ "$BUILD_DIR" == "frontend" ]; then
    # download and install general proto generator cli.
    wget -nv https://storage.googleapis.com/heurist-app.appspot.com/downloads/protoc
    chmod +x protoc
    sudo mv protoc /usr/local/bin/protoc

    if [ "$BUILD_DIR" == "backend" ]; then
        # download and install proto generator plugin for golang.
        wget -nv https://storage.googleapis.com/heurist-app.appspot.com/downloads/protoc-gen-go
        chmod +x protoc-gen-go
        sudo mv protoc-gen-go /usr/local/bin/protoc-gen-go
    fi

    if [ "$BUILD_DIR" == "frontend" ]; then
        # download and install proto generator plugin for grpc-web.
        wget -nv https://storage.googleapis.com/heurist-app.appspot.com/downloads/protoc-gen-grpc-web
        chmod +x protoc-gen-grpc-web
        sudo mv protoc-gen-grpc-web /usr/local/bin/protoc-gen-grpc-web
    fi

fi
