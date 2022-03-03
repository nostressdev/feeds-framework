package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) DeleteReaction(ctx context.Context, request *proto.DeleteReactionRequest) (*proto.DeleteReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	err := s.Storage.DeleteReaction(request.ReactionId)
	if err != nil {
		return nil, err
	}
	response := &proto.DeleteReactionResponse{
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
