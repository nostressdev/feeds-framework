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
	_, err := s.Storage.CreateGroupingFeed(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateGroupingFeedResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
