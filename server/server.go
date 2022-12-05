package server

import (
	"context"
	"log"

        "github.com/izaaklauer/mpotato/config"
	mpotatov1 "github.com/izaaklauer/mpotato/gen/proto/go/mpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MpotatoServer struct {
	mpotatov1.UnimplementedMpotatoServiceServer

	config config.Mpotato
}

// NewMpotatoServer initializes a new server from config
func NewMpotatoServer(config config.Mpotato) (*MpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := MpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * MpotatoServer) HelloWorld(
	ctx context.Context,
	req *mpotatov1.HelloWorldRequest,
) (*mpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &mpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
