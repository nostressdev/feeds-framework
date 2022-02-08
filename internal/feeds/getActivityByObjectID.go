package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetActivityByObjectID(ctx context.Context, request *proto.GetActivityByObjectIDRequest) (*proto.GetActivityByObjectIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	activity, err := s.Storage.GetActivityByObjectID(request.ForeignObjectId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetActivityByObjectIDResponse{
		Activity: activity,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
