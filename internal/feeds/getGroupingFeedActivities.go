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
	activity_groups, err := s.Storage.GetGroupingFeedActivities(request.GroupingFeedId, request.Limit, request.OffsetId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetGroupingFeedActivitiesResponse{
		ActivityGroups: activity_groups,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
