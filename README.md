# gRPC in the Browser

This project is created using the official [gRPC-Web](https://github.com/grpc/grpc-web)
library for connecting any gRPC api endpoint in any language.

gRPC-Web provides a Javascript library that lets browser clients access a gRPC
service. You can find out much more about gRPC in its own
[website](https://grpc.io).

## Overview

In this project I have used Go to create the backend, the idea here is to use the
firebase-admin-sdk to access Firestore in golang and I know that there is a
Firestore's client library for web that is already pretty fast but, I wanted to
create some small API's that has all access to the Firestore data and hide the data on client-side using the Firestore security rules.

Browsers only supports http protocol so, it is necessary to have a proxy that recieves the
`http://` and converts it to `grpc://` for backend to understand it. Envoy-proxy support grpc-web natively
so using some http-filters, so I have used the same.

## Guide

- The proto file is `heurist.proto` check the services defined there.
- Server's code is in Golang, you can use any other language.
- For `http://` to `grpc://` envoy-proxy is used.
- Client uses npm library grpc-web that gives the compatibility with the backend. And also Parcel bundler to bundle all the code.

## One Command Run

First you need to have a Firebase project ready and also Firestore initialized.
Then create some dummy data into your database.
Make sure to model the data in same way my backend service is using it.
Check out the service in `./backend/server.go`.

Get the firebase credentials.json for admin-sdk from your console.
In `./docker-compose.yml` copy the content of your JSON file add it like following

```yml
..
backend:
  image: rajveermalviya/grpc-browser:backend
  environment:
    FIREBASE_CREDENTIALS: '...YOUR...CREDENTIALS...JSON...CONTENT...' # replace this with actual content.
..
```

You also need to have [Docker](https://www.docker.com/get-started) & [Docker-Compose](https://docs.docker.com/compose/install/) installed in the dev-environment.

Then just run the following command from the repo root directory:

```sh
docker-compose up
```

Once the images are built, to see the _worst_ html page go to: `http://localhost:8080`

Now enter any username and then open the network panel it will send request to `https://localhost:8080/heurist_proto.HeuristGrpc` i.e the envoy-proxy will route it to grpc service.
