package main

import (
	"context"
	"flag"
	"io"
	"log"
	"fmt"
	"time"

	pb "github.com/adewinter/flockviz-server/routeguide"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func sendClickstream(client pb.RouteGuideClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clickPt := &pb.Point{Latitude: 2, Longitude: 2}
	numClicksToSend := 100

	stream, error := client.UserClickStream(ctx)
	if error != nil {
		log.Fatalf("Error with UserClickStream %v", error)
	}

	log.Println("Sending clickstream:")
	for i := 0; i < numClicksToSend; i++ {
		fmt.Print(".")
		error := stream.Send(clickPt)
		if error == io.EOF {
			break
		}
	}
	fmt.Println("Finished sending clicks")

	summary, error := stream.CloseAndRecv()
	if error != nil {
		log.Fatalf("Wow writing these print error handlers is boring %v", error)
	}
	log.Println("Summary: ", summary)
}

func getFlockTargetStream(client pb.RouteGuideClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	startPt := &pb.Point{Latitude: 1, Longitude: 1}
	startPtReq := &pb.FlockTargetStreamRequest{StartingLocation: startPt, TargetRatePerSecond: 2}
	stream, error := client.FlockTargetStream(ctx, startPtReq)

	if error != nil {
		log.Fatalf("Fatal error during getFlockTargetStream. Client: %v, error: %v", client, error)
	}

	for {
		point, error := stream.Recv()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatalf("Something went wrong during getFlockTargetStream while receiving a message. Client:%v, error: %v", client, error)
		}
		log.Println("HERE IS A POINT!:", point)
	}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	log.Println("Hello from client")
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	// getFlockTargetStream(client)
	sendClickstream(client)
}
