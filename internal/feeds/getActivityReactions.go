package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetActivityReactions(ctx context.Context, request *proto.GetActivityReactionsRequest) (*proto.GetActivityReactionsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	reactions, err := s.Storage.GetActivityReactions(request.ActivityId, request.Limit, request.OffsetId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetActivityReactionsResponse{
		Reactions: reactions,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
