package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetActivity(ctx context.Context, request *proto.GetActivityRequest) (*proto.GetActivityResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activity, err := s.Storage.GetActivity(request.ActivityId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetActivityResponse{
		Activity: activity,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
