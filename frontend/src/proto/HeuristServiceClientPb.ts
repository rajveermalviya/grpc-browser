/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';
import {
  CheckUsernameRequest,
  CheckUsernameResponse,
  GetUserDetailsRequest,
  GetUserDetailsResponse} from './heurist_pb';

export class HeuristGrpcClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCheck = new grpcWeb.AbstractClientBase.MethodInfo(
    CheckUsernameResponse,
    (request: CheckUsernameRequest) => {
      return request.serializeBinary();
    },
    CheckUsernameResponse.deserializeBinary
  );

  check(
    request: CheckUsernameRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: CheckUsernameResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/pb.HeuristGrpc/Check',
      request,
      metadata,
      this.methodInfoCheck,
      callback);
  }

  methodInfoGetUser = new grpcWeb.AbstractClientBase.MethodInfo(
    GetUserDetailsResponse,
    (request: GetUserDetailsRequest) => {
      return request.serializeBinary();
    },
    GetUserDetailsResponse.deserializeBinary
  );

  getUser(
    request: GetUserDetailsRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: GetUserDetailsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/pb.HeuristGrpc/GetUser',
      request,
      metadata,
      this.methodInfoGetUser,
      callback);
  }

}

