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
`http://` and converts it to `grpc://` for backend to understand it. So, official gRPC uses
Envoy proxy by default so I have used the same.

## Guide

- The proto file is `heurist.proto` check the services defined there.
- Server's code is in Golang, you can use any other language.
- For `http://` to `grpc://` envoy-proxy is used.
- Client uses npm library grpc-web that gives the compatibility with the backend. And also Parcel bundler to bundle all the code.

## One Command Run

First you need to have a Firebase project ready and also Firestore initialized.
Then create some dummy data into your database make sure to use the same schema as defined in the server.

You also need to have docker & docker-compose installed in the dev-environment.

Then just run the following command from the repo root directory:

```sh
docker-compose up
```

Once the images are built, and to see the worst html page go to:

`http://localhost:8080`

Now enter any username and then open the network panel it will send request to `localhost:8081` i.e the envoy-proxy.

If you have much better knowledge about nginx you can bypass this doing some rewriting of the hosts so that gRPC call can be
internally rewritten to envoy-proxy to have all the traffic going to only one host and only exposing one host.

To shutdown: `docker-compose down`.
