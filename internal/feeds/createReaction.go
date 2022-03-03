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
	reaction, err := s.Storage.CreateReaction(request.ActivityType, request.UserId, request.LinkedActivityId, request.Time, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateReactionResponse{
		Reaction: reaction,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
