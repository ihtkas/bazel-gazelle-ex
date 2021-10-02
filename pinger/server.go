package pinger

import (
	"context"

	pingerpb "github.com/ihtkas/bazel-gazelle-ex/api/pinger"
)

type Pinger struct {
	*pingerpb.UnimplementedPingerServiceServer
}

func (p *Pinger) Ping(ctx context.Context, req *pingerpb.PingRequest) (*pingerpb.PingResponse, error) {
	return &pingerpb.PingResponse{}, nil
}
