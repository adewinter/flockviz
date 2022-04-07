/**
 * @fileoverview gRPC-Web generated client stub for routeguide
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.routeguide = require('./route_guide_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.routeguide.RouteGuideClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.routeguide.RouteGuidePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.routeguide.FlockTargetStreamRequest,
 *   !proto.routeguide.Point>}
 */
const methodDescriptor_RouteGuide_FlockTargetStream = new grpc.web.MethodDescriptor(
  '/routeguide.RouteGuide/FlockTargetStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.routeguide.FlockTargetStreamRequest,
  proto.routeguide.Point,
  /**
   * @param {!proto.routeguide.FlockTargetStreamRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.routeguide.Point.deserializeBinary
);


/**
 * @param {!proto.routeguide.FlockTargetStreamRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.routeguide.Point>}
 *     The XHR Node Readable Stream
 */
proto.routeguide.RouteGuideClient.prototype.flockTargetStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/routeguide.RouteGuide/FlockTargetStream',
      request,
      metadata || {},
      methodDescriptor_RouteGuide_FlockTargetStream);
};


/**
 * @param {!proto.routeguide.FlockTargetStreamRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.routeguide.Point>}
 *     The XHR Node Readable Stream
 */
proto.routeguide.RouteGuidePromiseClient.prototype.flockTargetStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/routeguide.RouteGuide/FlockTargetStream',
      request,
      metadata || {},
      methodDescriptor_RouteGuide_FlockTargetStream);
};


module.exports = proto.routeguide;

