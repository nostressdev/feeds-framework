package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) AddExistingActivity(ctx context.Context, request *proto.AddExistingActivityRequest) (*proto.AddExistingActivityResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activity, err := s.Storage.AddExistingActivity(request.FeedId, request.ActivityId)
	if err != nil {
		return nil, err
	}
	response := &proto.AddExistingActivityResponse{
		Activity: activity,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
