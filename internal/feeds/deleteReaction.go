package feeds

import (
	"context"

	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/nerrors"
)

func (s *FeedsService) DeleteReaction(ctx context.Context, request *proto.DeleteReactionRequest) (*proto.DeleteReactionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	_, err := s.Storage.DeleteReaction(TODO)
	if err != nil {
		return nil, err
	}
	response := &proto.DeleteReactionResponse{
		TODO,
	}
	if err := response.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate response")
	}
	return response, nil
}
