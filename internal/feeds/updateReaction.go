package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) UpdateReaction(ctx context.Context, request *proto.UpdateReactionRequest) (*proto.UpdateReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	reaction, err := s.Storage.UpdateReaction(request.ReactionId, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.UpdateReactionResponse{
		Reaction: reaction,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
