package consultant

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetConsultantHandler struct {
	Model model.IModel
}

func (s *GetConsultantHandler) GetConsultant(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonConsultantResponse, error) {
	consultant, err := s.Model.GetConsultant(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ConsultantNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ConsultantToResponse(consultant)
	return resp, nil
}
