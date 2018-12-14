export class CheckUsernameRequest {
  constructor ();
  getUsername(): string;
  setUsername(a: string): void;
  toObject(): CheckUsernameRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CheckUsernameRequest;
}

export namespace CheckUsernameRequest {
  export type AsObject = {
    Username: string;
  }
}

export class CheckUsernameResponse {
  constructor ();
  getIsvalid(): boolean;
  setIsvalid(a: boolean): void;
  toObject(): CheckUsernameResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CheckUsernameResponse;
}

export namespace CheckUsernameResponse {
  export type AsObject = {
    Isvalid: boolean;
  }
}

export class GetUserDetailsRequest {
  constructor ();
  getUsername(): string;
  setUsername(a: string): void;
  toObject(): GetUserDetailsRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetUserDetailsRequest;
}

export namespace GetUserDetailsRequest {
  export type AsObject = {
    Username: string;
  }
}

export class GetUserDetailsResponse {
  constructor ();
  getUsername(): string;
  setUsername(a: string): void;
  getImageurl(): string;
  setImageurl(a: string): void;
  getName(): string;
  setName(a: string): void;
  getAbout(): string;
  setAbout(a: string): void;
  toObject(): GetUserDetailsResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetUserDetailsResponse;
}

export namespace GetUserDetailsResponse {
  export type AsObject = {
    Username: string;
    Imageurl: string;
    Name: string;
    About: string;
  }
}

