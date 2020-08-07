package question

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateQuestionHandler struct {
	Model model.IModel
}

func (s *UpdateQuestionHandler) UpdateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	question := s.reqToQuestion(req)

	v, err := s.Model.UpdateQuestion(ctx, question)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.QuestionNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.QuestionToResponse(v)
	return resp, nil
}

func (s *UpdateQuestionHandler) reqToQuestion(req *pb.CommonQuestionRequest) *dto.Question {
	question := &dto.Question{
		ID:       utility.RemoveZeroWidth(req.Id),
		Category: req.Data.Category,
		Type:     req.Data.Type,
		Content:  req.Data.Content,
	}
	return question
}
