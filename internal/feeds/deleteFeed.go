package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) DeleteFeed(ctx context.Context, request *proto.DeleteFeedRequest) (*proto.DeleteFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	err := s.Storage.DeleteFeed(request.FeedId)
	if err != nil {
		return nil, err
	}
	response := &proto.DeleteFeedResponse{
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
