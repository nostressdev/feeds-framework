package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) AddReaction(ctx context.Context, request *proto.AddReactionRequest) (*proto.AddReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	err := s.Storage.AddReaction(request.FeedId, request.ReactionId)
	if err != nil {
		return nil, err
	}
	response := &proto.AddReactionResponse{
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
