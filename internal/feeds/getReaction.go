package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetReaction(ctx context.Context, request *proto.GetReactionRequest) (*proto.GetReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	reaction, err := s.Storage.GetReaction(request.ReactionId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetReactionResponse{
		Reaction: reaction,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
