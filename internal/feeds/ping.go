package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
)

func (s *FeedsService) Ping(ctx context.Context, request *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{}, nil
}
