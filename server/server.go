package main

import (
	"flag"
	"fmt"
	pb "github.com/adewinter/flockviz-server/routeguide"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (rgServer *routeGuideServer) UserClickStream(clickStream pb.RouteGuide_UserClickStreamServer) error {
	fmt.Println("GOT A RPC REQUEST FOR USER_CLICK_STREAM")
	for {
		data, error := clickStream.Recv()
		if error == io.EOF {
			clickStream.SendAndClose(&pb.ClickSummary{Status: 12})
			break
		}
		if error != nil {
			log.Fatalf("Got some kind of error in UserClickStream:%v", error)
		}
		fmt.Print("data:", data)

	}
	return nil
}

func (rgServer *routeGuideServer) FlockTargetStream(streamRequestMessage *pb.FlockTargetStreamRequest, targetStream pb.RouteGuide_FlockTargetStreamServer) error {
	fmt.Println("GOT A RPC REQUEST FOR FLOCK_TARGET_STREAM")

	// send a single point via the stream, could do this in a for loop for multiple points
	pointsToSend := 100
	fmt.Println("Sending Points:")
	for i := 0; i < pointsToSend; i++ {
		point := new(pb.Point)
		point.Latitude = 1
		point.Longitude = 1
		fmt.Print(".")
		targetStream.Send(point)
	}
	return nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-grpc-web, *")
}

func main() {
	flag.Parse()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	foo := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		setupCorsResponse(&resp, req)
		if (*req).Method == "OPTIONS" {
			return
		}
		enableCors(&resp)
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}
		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	log.Println("Serving on https://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", foo))

}
