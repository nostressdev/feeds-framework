package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetFeed(ctx context.Context, request *proto.GetFeedRequest) (*proto.GetFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	feed, err := s.Storage.GetFeed(request.FeedId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetFeedResponse{
		Feed: feed,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
