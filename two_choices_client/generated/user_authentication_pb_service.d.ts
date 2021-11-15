// package: two_choices
// file: user_authentication.proto

import * as user_authentication_pb from "./user_authentication_pb";
import {grpc} from "@improbable-eng/grpc-web";

type UserAuthenticationServiceSendVerificationMail = {
  readonly methodName: string;
  readonly service: typeof UserAuthenticationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_authentication_pb.UserAuthentication.SendVerificationMailRequest;
  readonly responseType: typeof user_authentication_pb.UserAuthentication.SendVerificationMailResponse;
};

type UserAuthenticationServiceUserRegistration = {
  readonly methodName: string;
  readonly service: typeof UserAuthenticationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_authentication_pb.UserAuthentication.UserRegistrationRequest;
  readonly responseType: typeof user_authentication_pb.UserAuthentication.UserRegistrationResponse;
};

export class UserAuthenticationService {
  static readonly serviceName: string;
  static readonly SendVerificationMail: UserAuthenticationServiceSendVerificationMail;
  static readonly UserRegistration: UserAuthenticationServiceUserRegistration;
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

export class UserAuthenticationServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  sendVerificationMail(
    requestMessage: user_authentication_pb.UserAuthentication.SendVerificationMailRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_authentication_pb.UserAuthentication.SendVerificationMailResponse|null) => void
  ): UnaryResponse;
  sendVerificationMail(
    requestMessage: user_authentication_pb.UserAuthentication.SendVerificationMailRequest,
    callback: (error: ServiceError|null, responseMessage: user_authentication_pb.UserAuthentication.SendVerificationMailResponse|null) => void
  ): UnaryResponse;
  userRegistration(
    requestMessage: user_authentication_pb.UserAuthentication.UserRegistrationRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_authentication_pb.UserAuthentication.UserRegistrationResponse|null) => void
  ): UnaryResponse;
  userRegistration(
    requestMessage: user_authentication_pb.UserAuthentication.UserRegistrationRequest,
    callback: (error: ServiceError|null, responseMessage: user_authentication_pb.UserAuthentication.UserRegistrationResponse|null) => void
  ): UnaryResponse;
}

