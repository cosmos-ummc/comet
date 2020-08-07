package question

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateQuestionHandler struct {
	Model model.IModel
}

func (s *CreateQuestionHandler) CreateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	question := &dto.Question{
		ID:       uuid.NewV4().String(),
		Category: req.Data.Category,
		Type:     req.Data.Type,
		Content:  req.Data.Content,
		Score:    req.Data.Score,
	}

	rslt, err := s.Model.CreateQuestion(ctx, question)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.QuestionAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.QuestionToResponse(rslt)
	return resp, nil
}
