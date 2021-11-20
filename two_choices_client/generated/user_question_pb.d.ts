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
      title: string,
      firstAnswer: string,
      secondAnswer: string,
      firstCount: number,
      secondCount: number,
    }
  }
}

