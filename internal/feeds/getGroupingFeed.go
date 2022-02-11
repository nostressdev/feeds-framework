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
	_, err := s.Storage.GetGroupingFeed(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.GetGroupingFeedResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
