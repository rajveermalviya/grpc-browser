export class CheckUsernameRequest {
  constructor();
  getUsername(): string;
  setUsername(a: string): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CheckUsernameRequest;
}

export class CheckUsernameResponse {
  constructor();
  getIsvalid(): boolean;
  setIsvalid(a: boolean): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => CheckUsernameResponse;
}

export class GetUserDetailsRequest {
  constructor();
  getUsername(): string;
  setUsername(a: string): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetUserDetailsRequest;
}

export class GetUserDetailsResponse {
  constructor();
  getUsername(): string;
  setUsername(a: string): void;
  getImageurl(): string;
  setImageurl(a: string): void;
  getName(): string;
  setName(a: string): void;
  getAbout(): string;
  setAbout(a: string): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetUserDetailsResponse;
}
