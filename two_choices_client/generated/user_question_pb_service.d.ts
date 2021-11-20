// package: two_choices
// file: user_question.proto

import * as user_question_pb from "./user_question_pb";
import * as common_pb from "./common_pb";
import {grpc} from "@improbable-eng/grpc-web";

type UserQuestionServiceGetRandom = {
  readonly methodName: string;
  readonly service: typeof UserQuestionService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof common_pb.Empty;
  readonly responseType: typeof user_question_pb.UserQuestion.GetRandomResponse;
};

export class UserQuestionService {
  static readonly serviceName: string;
  static readonly GetRandom: UserQuestionServiceGetRandom;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class UserQuestionServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getRandom(
    requestMessage: common_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_question_pb.UserQuestion.GetRandomResponse|null) => void
  ): UnaryResponse;
  getRandom(
    requestMessage: common_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: user_question_pb.UserQuestion.GetRandomResponse|null) => void
  ): UnaryResponse;
}

