// package: two_choices
// file: user_question.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class UserQuestion extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserQuestion.AsObject;
  static toObject(includeInstance: boolean, msg: UserQuestion): UserQuestion.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserQuestion, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserQuestion;
  static deserializeBinaryFromReader(message: UserQuestion, reader: jspb.BinaryReader): UserQuestion;
}

export namespace UserQuestion {
  export type AsObject = {
  }

  export class GetRandomResponse extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    getTitle(): string;
    setTitle(value: string): void;

    getFirstAnswer(): string;
    setFirstAnswer(value: string): void;

    getSecondAnswer(): string;
    setSecondAnswer(value: string): void;

    getFirstCount(): number;
    setFirstCount(value: number): void;

    getSecondCount(): number;
    setSecondCount(value: number): void;

    getFirstImgUrl(): string;
    setFirstImgUrl(value: string): void;

    getSecondImgUrl(): string;
    setSecondImgUrl(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRandomResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetRandomResponse): GetRandomResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRandomResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRandomResponse;
    static deserializeBinaryFromReader(message: GetRandomResponse, reader: jspb.BinaryReader): GetRandomResponse;
  }

  export namespace GetRandomResponse {
    export type AsObject = {
      id: number,
      title: string,
      firstAnswer: string,
      secondAnswer: string,
      firstCount: number,
      secondCount: number,
      firstImgUrl: string,
      secondImgUrl: string,
    }
  }

  export class UpdateRequest extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    getIsFirstSelected(): boolean;
    setIsFirstSelected(value: boolean): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequest;
    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
  }

  export namespace UpdateRequest {
    export type AsObject = {
      id: number,
      isFirstSelected: boolean,
    }
  }
}

