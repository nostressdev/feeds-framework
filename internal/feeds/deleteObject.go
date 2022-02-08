package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) DeleteObject(ctx context.Context, request *proto.DeleteObjectRequest) (*proto.DeleteObjectResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	err := s.Storage.DeleteObject(request.CollectionId, request.ObjectId)
	if err != nil {
		return nil, err
	}
	response := &proto.DeleteObjectResponse{
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
