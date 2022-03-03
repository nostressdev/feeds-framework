package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) UpdateFeed(ctx context.Context, request *proto.UpdateFeedRequest) (*proto.UpdateFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	feed, err := s.Storage.UpdateFeed(request.FeedId, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.UpdateFeedResponse{
		Feed: feed,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
