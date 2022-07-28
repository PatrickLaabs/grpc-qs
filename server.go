/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/PatrickLaabs/grpc-qs/healthchecker"

	//_ "github.com/PatrickLaabs/grpc-qs/healthchecker"
	"log"
	"net"

	pb "github.com/PatrickLaabs/grpc-qs/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedByeServer
	pb.UnimplementedVersionServer
	pb.UnimplementedEitcoServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayBye implements helloworld.ByeServer
func (s *server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ByeReply{Message: "Bye " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

// SendVersion func
func (s *server) SendVersion(ctx context.Context, in *pb.VersionRequest) (*pb.VersionReply, error) {
	f, err := os.ReadFile("./serverVersion.params")
	if err != nil {
		log.Fatal("could not open param file", err)
	}
	return &pb.VersionReply{Message: "Current gRPC Server Version: " + string(f)}, nil
}

// SayEitco
func (s *server) SayEitco(ctx context.Context, in *pb.HelloEitcoRequest) (*pb.HelloEitcoReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloEitcoReply{Message: "Eitco " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterByeServer(s, &server{})
	pb.RegisterVersionServer(s, &server{})
	pb.RegisterEitcoServer(s, &server{})

	healthService := healthchecker.HealthChecker{}
	grpc_health_v1.RegisterHealthServer(s, &healthService)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
