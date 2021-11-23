// package: two_choices
// file: user_question.proto

var user_question_pb = require("./user_question_pb");
var common_pb = require("./common_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UserQuestionService = (function () {
  function UserQuestionService() {}
  UserQuestionService.serviceName = "two_choices.UserQuestionService";
  return UserQuestionService;
}());

UserQuestionService.GetRandom = {
  methodName: "GetRandom",
  service: UserQuestionService,
  requestStream: false,
  responseStream: false,
  requestType: common_pb.Empty,
  responseType: user_question_pb.UserQuestion.GetRandomResponse
};

UserQuestionService.Update = {
  methodName: "Update",
  service: UserQuestionService,
  requestStream: false,
  responseStream: false,
  requestType: user_question_pb.UserQuestion.UpdateRequest,
  responseType: common_pb.Empty
};

exports.UserQuestionService = UserQuestionService;

function UserQuestionServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserQuestionServiceClient.prototype.getRandom = function getRandom(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserQuestionService.GetRandom, {
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

UserQuestionServiceClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserQuestionService.Update, {
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

exports.UserQuestionServiceClient = UserQuestionServiceClient;

