package consultant

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
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
	
	// delete user
	_, users, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.ConsultantID,
		Value: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(users) != 0 {
		for _, user := range users {
			_, err = s.Model.DeleteUser(ctx, user.ID)
			if err != nil {
				continue
			}
		}
	}
	
	resp := utility.ConsultantToResponse(rslt)
	return resp, nil
}
