package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) AddActivity(ctx context.Context, request *proto.AddActivityRequest) (*proto.AddActivityResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activity, err := s.Storage.AddActivity(request.FeedId, request.ForeignObjectId, request.UserId, request.ActivityType, request.Time, request.RedirectTo, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.AddActivityResponse{
		Activity: activity,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
