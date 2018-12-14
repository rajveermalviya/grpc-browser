import {
  CheckUsernameRequest,
  GetUserDetailsRequest,
  CheckUsernameResponse,
  GetUserDetailsResponse
} from "./proto/heurist_pb";
import { HeuristGrpcClient } from "./proto/HeuristServiceClientPb";
import { Metadata } from "grpc-web";

export class GRPC {
  private _client: HeuristGrpcClient;
  constructor(client: HeuristGrpcClient) {
    this._client = client;
  }
  check(req: CheckUsernameRequest, metadata: Metadata) {
    return new Promise<CheckUsernameResponse>((resolve, reject) => {
      this._client.check(req, metadata, (err, res) => {
        if (err) reject(err);
        resolve(res);
      });
    });
  }
  getUser(req: GetUserDetailsRequest, metadata: Metadata) {
    return new Promise<GetUserDetailsResponse>((resolve, reject) => {
      this._client.getUser(req, metadata, (err, res) => {
        if (err) reject(err);
        resolve(res);
      });
    });
  }
}
