package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) UpdateActivity(ctx context.Context, request *proto.UpdateActivityRequest) (*proto.UpdateActivityResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activity, err := s.Storage.UpdateActivity(request.ActivityId, request.ExtraData)
	if err != nil {
		return nil, err
	}
	response := &proto.UpdateActivityResponse{
		Activity: activity,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
