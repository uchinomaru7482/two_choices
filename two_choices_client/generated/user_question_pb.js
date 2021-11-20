// source: user_question.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var common_pb = require('./common_pb.js');
goog.object.extend(proto, common_pb);
goog.exportSymbol('proto.two_choices.UserQuestion', null, global);
goog.exportSymbol('proto.two_choices.UserQuestion.GetRandomResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.two_choices.UserQuestion = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.two_choices.UserQuestion, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.two_choices.UserQuestion.displayName = 'proto.two_choices.UserQuestion';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.two_choices.UserQuestion.GetRandomResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.two_choices.UserQuestion.GetRandomResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.two_choices.UserQuestion.GetRandomResponse.displayName = 'proto.two_choices.UserQuestion.GetRandomResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.two_choices.UserQuestion.prototype.toObject = function(opt_includeInstance) {
  return proto.two_choices.UserQuestion.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.two_choices.UserQuestion} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.two_choices.UserQuestion.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.two_choices.UserQuestion}
 */
proto.two_choices.UserQuestion.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.two_choices.UserQuestion;
  return proto.two_choices.UserQuestion.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.two_choices.UserQuestion} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.two_choices.UserQuestion}
 */
proto.two_choices.UserQuestion.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.two_choices.UserQuestion.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.two_choices.UserQuestion.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.two_choices.UserQuestion} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.two_choices.UserQuestion.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.two_choices.UserQuestion.GetRandomResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.two_choices.UserQuestion.GetRandomResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.two_choices.UserQuestion.GetRandomResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    title: jspb.Message.getFieldWithDefault(msg, 1, ""),
    firstAnswer: jspb.Message.getFieldWithDefault(msg, 2, ""),
    secondAnswer: jspb.Message.getFieldWithDefault(msg, 3, ""),
    firstCount: jspb.Message.getFieldWithDefault(msg, 4, 0),
    secondCount: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse}
 */
proto.two_choices.UserQuestion.GetRandomResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.two_choices.UserQuestion.GetRandomResponse;
  return proto.two_choices.UserQuestion.GetRandomResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.two_choices.UserQuestion.GetRandomResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse}
 */
proto.two_choices.UserQuestion.GetRandomResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFirstAnswer(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSecondAnswer(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setFirstCount(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setSecondCount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.two_choices.UserQuestion.GetRandomResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.two_choices.UserQuestion.GetRandomResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.two_choices.UserQuestion.GetRandomResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFirstAnswer();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSecondAnswer();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFirstCount();
  if (f !== 0) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = message.getSecondCount();
  if (f !== 0) {
    writer.writeUint64(
      5,
      f
    );
  }
};


/**
 * optional string title = 1;
 * @return {string}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse} returns this
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.setTitle = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string first_answer = 2;
 * @return {string}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.getFirstAnswer = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse} returns this
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.setFirstAnswer = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string second_answer = 3;
 * @return {string}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.getSecondAnswer = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse} returns this
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.setSecondAnswer = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional uint64 first_count = 4;
 * @return {number}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.getFirstCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse} returns this
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.setFirstCount = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional uint64 second_count = 5;
 * @return {number}
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.getSecondCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.two_choices.UserQuestion.GetRandomResponse} returns this
 */
proto.two_choices.UserQuestion.GetRandomResponse.prototype.setSecondCount = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


goog.object.extend(exports, proto.two_choices);
