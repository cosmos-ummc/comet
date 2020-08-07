package question

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateQuestionsHandler struct {
	Model model.IModel
}

func (s *UpdateQuestionsHandler) UpdateQuestions(ctx context.Context, req *pb.CommonQuestionsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	question := s.reqToQuestion(req)

	ids, err := s.Model.UpdateQuestions(ctx, question, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.QuestionNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateQuestionsHandler) reqToQuestion(req *pb.CommonQuestionsRequest) *dto.Question {
	question := &dto.Question{
		Category: req.Data.Category,
		Type:     req.Data.Type,
		Content:  req.Data.Content,
		ID:       req.Ids[0],
	}
	return question
}
