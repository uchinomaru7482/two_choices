// package: two_choices
// file: user_cert_profile.proto

import * as jspb from "google-protobuf";

export class UserCertProfile extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCertProfile.AsObject;
  static toObject(includeInstance: boolean, msg: UserCertProfile): UserCertProfile.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserCertProfile, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCertProfile;
  static deserializeBinaryFromReader(message: UserCertProfile, reader: jspb.BinaryReader): UserCertProfile;
}

export namespace UserCertProfile {
  export type AsObject = {
  }

  export class GetProfileRequest extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetProfileRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetProfileRequest): GetProfileRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetProfileRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetProfileRequest;
    static deserializeBinaryFromReader(message: GetProfileRequest, reader: jspb.BinaryReader): GetProfileRequest;
  }

  export namespace GetProfileRequest {
    export type AsObject = {
      id: number,
    }
  }

  export class GetProfileResponse extends jspb.Message {
    getName(): string;
    setName(value: string): void;

    getEmail(): string;
    setEmail(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetProfileResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetProfileResponse): GetProfileResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetProfileResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetProfileResponse;
    static deserializeBinaryFromReader(message: GetProfileResponse, reader: jspb.BinaryReader): GetProfileResponse;
  }

  export namespace GetProfileResponse {
    export type AsObject = {
      name: string,
      email: string,
    }
  }
}

