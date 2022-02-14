package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) CreateGroupingFeed(ctx context.Context, request *proto.CreateGroupingFeedRequest) (*proto.CreateGroupingFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	feed, err := s.Storage.CreateGroupingFeed(request.UserId, request.KeyFormat, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateGroupingFeedResponse{
		GroupingFeed: feed,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
