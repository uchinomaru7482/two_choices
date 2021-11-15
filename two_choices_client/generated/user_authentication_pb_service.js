// package: two_choices
// file: user_authentication.proto

var user_authentication_pb = require("./user_authentication_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UserAuthenticationService = (function () {
  function UserAuthenticationService() {}
  UserAuthenticationService.serviceName = "two_choices.UserAuthenticationService";
  return UserAuthenticationService;
}());

UserAuthenticationService.SendVerificationMail = {
  methodName: "SendVerificationMail",
  service: UserAuthenticationService,
  requestStream: false,
  responseStream: false,
  requestType: user_authentication_pb.UserAuthentication.SendVerificationMailRequest,
  responseType: user_authentication_pb.UserAuthentication.SendVerificationMailResponse
};

UserAuthenticationService.UserRegistration = {
  methodName: "UserRegistration",
  service: UserAuthenticationService,
  requestStream: false,
  responseStream: false,
  requestType: user_authentication_pb.UserAuthentication.UserRegistrationRequest,
  responseType: user_authentication_pb.UserAuthentication.UserRegistrationResponse
};

exports.UserAuthenticationService = UserAuthenticationService;

function UserAuthenticationServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserAuthenticationServiceClient.prototype.sendVerificationMail = function sendVerificationMail(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserAuthenticationService.SendVerificationMail, {
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

UserAuthenticationServiceClient.prototype.userRegistration = function userRegistration(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserAuthenticationService.UserRegistration, {
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

exports.UserAuthenticationServiceClient = UserAuthenticationServiceClient;

