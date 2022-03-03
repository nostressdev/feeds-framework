package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) GetObject(ctx context.Context, request *proto.GetObjectRequest) (*proto.GetObjectResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	object, err := s.Storage.GetObject(request.CollectionId, request.ObjectId)
	if err != nil {
		return nil, err
	}
	response := &proto.GetObjectResponse{
		Object: object,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
