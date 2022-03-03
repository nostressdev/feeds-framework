package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) CreateObject(ctx context.Context, request *proto.CreateObjectRequest) (*proto.CreateObjectResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	object, err := s.Storage.CreateObject(request.CollectionId, request.ObjectId, request.Data)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateObjectResponse{
		Object: object,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
