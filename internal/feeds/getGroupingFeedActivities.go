package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetGroupingFeedActivities(ctx context.Context, request *proto.GetGroupingFeedActivitiesRequest) (*proto.GetGroupingFeedActivitiesResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	_, err := s.Storage.GetGroupingFeedActivities(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.GetGroupingFeedActivitiesResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
