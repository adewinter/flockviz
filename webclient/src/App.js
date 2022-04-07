import React, { Component } from "react";
import { useState } from "react";

import "./App.css";

import { RouteGuideClient } from "./routeguide/route_guide_grpc_web_pb";
import {
  Point,
  ClickSummary,
  FlockTargetStreamRequest,
} from "./routeguide/route_guide_pb";

function Welcome(props) {
  return <h1>WELCOME, {props.name}, to Flockviz!</h1>;
}

function pointToString(point) {
  return point.getLatitude() + ", " + point.getLongitude();
}

function PointComponent(props) {
  return <span>({pointToString(props.point)})&nbsp;</span>;
}

function App() {
  const [components, setState] = useState({ points: [] });
  const routeGuideService = new RouteGuideClient("http://localhost:10000");
  const flockRequest = new FlockTargetStreamRequest();
  flockRequest.setTargetratepersecond(2);

  function addPointToState(point) {
    setState((prevState) => ({
      points: [...prevState.points, point],
    }));
  }

  function clearState() {
    setState((prevState) => ({
      points: [],
    }));
  }

  function getPoints() {
    console.log("Requesting FlockTargetStream");
    clearState();

    var ftStream = routeGuideService.flockTargetStream(flockRequest);
    var counter = 0;

    ftStream.on("data", function (point) {
      console.log("Got point:", pointToString(point));
      addPointToState(point);
    });

    ftStream.on("end", function (end) {
      console.log("This is the end!", end);
    });
  }

  return (
    <div className="App">
      <Welcome name="SOMEONE" />
      <button onClick={getPoints}>Go</button>
      <div>
        {components["points"].map(function (point, i) {
          return <PointComponent point={point} key={i} />;
        })}
      </div>
    </div>
  );
}

export default App;
