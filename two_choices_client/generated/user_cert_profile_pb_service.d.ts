// package: two_choices
// file: user_cert_profile.proto

import * as user_cert_profile_pb from "./user_cert_profile_pb";
import {grpc} from "@improbable-eng/grpc-web";

type UserCertProfileServiceGetProfile = {
  readonly methodName: string;
  readonly service: typeof UserCertProfileService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_cert_profile_pb.UserCertProfile.GetProfileRequest;
  readonly responseType: typeof user_cert_profile_pb.UserCertProfile.GetProfileResponse;
};

export class UserCertProfileService {
  static readonly serviceName: string;
  static readonly GetProfile: UserCertProfileServiceGetProfile;
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

export class UserCertProfileServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getProfile(
    requestMessage: user_cert_profile_pb.UserCertProfile.GetProfileRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_cert_profile_pb.UserCertProfile.GetProfileResponse|null) => void
  ): UnaryResponse;
  getProfile(
    requestMessage: user_cert_profile_pb.UserCertProfile.GetProfileRequest,
    callback: (error: ServiceError|null, responseMessage: user_cert_profile_pb.UserCertProfile.GetProfileResponse|null) => void
  ): UnaryResponse;
}

