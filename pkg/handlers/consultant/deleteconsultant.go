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

type DeleteConsultantHandler struct {
	Model model.IModel
}

func (s *DeleteConsultantHandler) DeleteConsultant(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonConsultantResponse, error) {
	rslt, err := s.Model.DeleteConsultant(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ConsultantNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ConsultantToResponse(rslt)
	return resp, nil
}
