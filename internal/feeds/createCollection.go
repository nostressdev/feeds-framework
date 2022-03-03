package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) CreateCollection(ctx context.Context, request *proto.CreateCollectionRequest) (*proto.CreateCollectionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	collection, err := s.Storage.CreateCollection(request.CollectionId, request.DeletingType)
	if err != nil {
		return nil, err
	}
	response := &proto.CreateCollectionResponse{
		Collection: collection,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
