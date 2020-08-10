package consultant

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type DeleteConsultantHandler struct {
	Model model.IModel
}

func (s *DeleteConsultantHandler) DeleteConsultant(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonConsultantResponse, error) {
	rslt, err := s.Model.DeleteConsultant(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	
	resp := utility.ConsultantToResponse(rslt)
	return resp, nil
}
