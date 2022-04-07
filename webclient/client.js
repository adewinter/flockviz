const {
  Point,
  ClickSummary,
  FlockTargetStreamRequest,
} = require("./routeguide/route_guide_pb.js");
const { RouteGuideClient } = require("./routeguide/route_guide_grpc_web_pb.js");

console.log("HELLO!!");
var routeGuideService = new RouteGuideClient("http://localhost:10000");
console.log("This is the service", routeGuideService);
var flockRequest = new FlockTargetStreamRequest();
flockRequest.setTargetratepersecond(2);

var ftStream = routeGuideService.flockTargetStream(flockRequest);

ftStream.on("data", function (point) {
  console.log("This is a point?", point);
  console.log("Lat:", point.getLatitude());
});

ftStream.on("end", function (end) {
  console.log("This is the end!", end);
});
