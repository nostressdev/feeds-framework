package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) CreateReaction(ctx context.Context, request *proto.CreateReactionRequest) (*proto.CreateReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	_, err := s.Storage.CreateReaction(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateReactionResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}