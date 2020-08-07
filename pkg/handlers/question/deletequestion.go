package question

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteQuestionHandler struct {
	Model model.IModel
}

func (s *DeleteQuestionHandler) DeleteQuestion(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonQuestionResponse, error) {
	rslt, err := s.Model.DeleteQuestion(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.QuestionNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.QuestionToResponse(rslt)
	return resp, nil
}
