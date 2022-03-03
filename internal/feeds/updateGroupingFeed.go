package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) UpdateGroupingFeed(ctx context.Context, request *proto.UpdateGroupingFeedRequest) (*proto.UpdateGroupingFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	feed, err := s.Storage.UpdateGroupingFeed(request.GroupingFeedId, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.UpdateGroupingFeedResponse{
		GroupingFeed: feed,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
