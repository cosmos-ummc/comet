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

type GetQuestionHandler struct {
	Model model.IModel
}

func (s *GetQuestionHandler) GetQuestion(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonQuestionResponse, error) {
	question, err := s.Model.GetQuestion(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.QuestionNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.QuestionToResponse(question)
	return resp, nil
}
