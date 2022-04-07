package main

import (
	"flag"
	"fmt"
	"log"
	"io"
	"net/http"
	// "time"

	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/examples/data"

	// "github.com/golang/protobuf/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	pb "github.com/adewinter/flockviz-server/routeguide"
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
	for i := 0; i < pointsToSend; i++ {
		point := new(pb.Point)
		point.Latitude = 1
		point.Longitude = 1
		targetStream.Send(point)
	}
	return nil
}

func main() {
	flag.Parse()
	// fmt.Println("Server start on Port:", *port)
	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	// if err != nil {
	// 	log.Fatalf("Error opening listener. Err: %v", err)
	// }

	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	// httpSrv := &http.Server{
	// 	// These interfere with websocket streams, disable for now
	// 	// ReadTimeout: 5 * time.Second,
	// 	// WriteTimeout: 10 * time.Second,
	// 	ReadHeaderTimeout: 5 * time.Second,
	// 	IdleTimeout:       120 * time.Second,
	// 	Addr:              ":https",
	// 	// TLSConfig: &tls.Config{
	// 	// 	PreferServerCipherSuites: true,
	// 	// 	CurvePreferences: []tls.CurveID{
	// 	// 		tls.CurveP256,
	// 	// 		tls.X25519,
	// 	// 	},
	// 	// },
	// 	Handler: ,
	// }

	foo := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Here?", req)
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}	
		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	log.Println("Serving on https://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", foo))

	
	// grpcServer.Serve(lis)
}
