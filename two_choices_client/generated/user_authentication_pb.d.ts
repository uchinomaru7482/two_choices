// package: two_choices
// file: user_authentication.proto

import * as jspb from "google-protobuf";

export class UserAuthentication extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserAuthentication.AsObject;
  static toObject(includeInstance: boolean, msg: UserAuthentication): UserAuthentication.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserAuthentication, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserAuthentication;
  static deserializeBinaryFromReader(message: UserAuthentication, reader: jspb.BinaryReader): UserAuthentication;
}

export namespace UserAuthentication {
  export type AsObject = {
  }

  export class SendVerificationMailRequest extends jspb.Message {
    getName(): string;
    setName(value: string): void;

    getEmail(): string;
    setEmail(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SendVerificationMailRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SendVerificationMailRequest): SendVerificationMailRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SendVerificationMailRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SendVerificationMailRequest;
    static deserializeBinaryFromReader(message: SendVerificationMailRequest, reader: jspb.BinaryReader): SendVerificationMailRequest;
  }

  export namespace SendVerificationMailRequest {
    export type AsObject = {
      name: string,
      email: string,
    }
  }

  export class SendVerificationMailResponse extends jspb.Message {
    getResult(): UserAuthentication.SendVerificationMailResultMap[keyof UserAuthentication.SendVerificationMailResultMap];
    setResult(value: UserAuthentication.SendVerificationMailResultMap[keyof UserAuthentication.SendVerificationMailResultMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SendVerificationMailResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SendVerificationMailResponse): SendVerificationMailResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SendVerificationMailResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SendVerificationMailResponse;
    static deserializeBinaryFromReader(message: SendVerificationMailResponse, reader: jspb.BinaryReader): SendVerificationMailResponse;
  }

  export namespace SendVerificationMailResponse {
    export type AsObject = {
      result: UserAuthentication.SendVerificationMailResultMap[keyof UserAuthentication.SendVerificationMailResultMap],
    }
  }

  export class UserRegistrationRequest extends jspb.Message {
    getState(): string;
    setState(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserRegistrationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UserRegistrationRequest): UserRegistrationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserRegistrationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserRegistrationRequest;
    static deserializeBinaryFromReader(message: UserRegistrationRequest, reader: jspb.BinaryReader): UserRegistrationRequest;
  }

  export namespace UserRegistrationRequest {
    export type AsObject = {
      state: string,
    }
  }

  export class UserRegistrationResponse extends jspb.Message {
    getResult(): UserAuthentication.UserRegistrationResultMap[keyof UserAuthentication.UserRegistrationResultMap];
    setResult(value: UserAuthentication.UserRegistrationResultMap[keyof UserAuthentication.UserRegistrationResultMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserRegistrationResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UserRegistrationResponse): UserRegistrationResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserRegistrationResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserRegistrationResponse;
    static deserializeBinaryFromReader(message: UserRegistrationResponse, reader: jspb.BinaryReader): UserRegistrationResponse;
  }

  export namespace UserRegistrationResponse {
    export type AsObject = {
      result: UserAuthentication.UserRegistrationResultMap[keyof UserAuthentication.UserRegistrationResultMap],
    }
  }

  export interface SendVerificationMailResultMap {
    SEND_MAIL_SUCCESS: 0;
    ALREADY_EXISTS_ADDRESS: 1;
  }

  export const SendVerificationMailResult: SendVerificationMailResultMap;

  export interface UserRegistrationResultMap {
    USER_REGISTRATION_SUCCESS: 0;
    TIMEOUT: 1;
    INCORRECT_STATE: 2;
  }

  export const UserRegistrationResult: UserRegistrationResultMap;
}

