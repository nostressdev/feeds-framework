package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetFeedActivities(ctx context.Context, request *proto.GetFeedActivitiesRequest) (*proto.GetFeedActivitiesResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activities, err := s.Storage.GetFeedActivities(request.FeedId, request.Limit, request.OffsetId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetFeedActivitiesResponse{
		Activities: activities,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
