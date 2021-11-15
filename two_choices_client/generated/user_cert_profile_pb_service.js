// package: two_choices
// file: user_cert_profile.proto

var user_cert_profile_pb = require("./user_cert_profile_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UserCertProfileService = (function () {
  function UserCertProfileService() {}
  UserCertProfileService.serviceName = "two_choices.UserCertProfileService";
  return UserCertProfileService;
}());

UserCertProfileService.GetProfile = {
  methodName: "GetProfile",
  service: UserCertProfileService,
  requestStream: false,
  responseStream: false,
  requestType: user_cert_profile_pb.UserCertProfile.GetProfileRequest,
  responseType: user_cert_profile_pb.UserCertProfile.GetProfileResponse
};

exports.UserCertProfileService = UserCertProfileService;

function UserCertProfileServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserCertProfileServiceClient.prototype.getProfile = function getProfile(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserCertProfileService.GetProfile, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.UserCertProfileServiceClient = UserCertProfileServiceClient;

