package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetGroupingFeed(ctx context.Context, request *proto.GetGroupingFeedRequest) (*proto.GetGroupingFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	feed, err := s.Storage.GetGroupingFeed(request.GroupingFeedId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetGroupingFeedResponse{
		GroupingFeed: feed,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
