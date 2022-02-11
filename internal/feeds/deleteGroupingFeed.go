package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) DeleteGroupingFeed(ctx context.Context, request *proto.DeleteGroupingFeedRequest) (*proto.DeleteGroupingFeedResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	_, err := s.Storage.DeleteGroupingFeed(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.DeleteGroupingFeedResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
