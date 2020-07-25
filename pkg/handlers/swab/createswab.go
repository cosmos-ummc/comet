package swab

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

type CreateSwabHandler struct {
	Model model.IModel
}

func (s *CreateSwabHandler) CreateSwab(ctx context.Context, req *pb.CommonSwabRequest, user *dto.User) (*pb.CommonSwabResponse, error) {
	swab := &dto.Swab{
		PatientID:           utility.RemoveZeroWidth(req.Data.PatientId),
		Status:              req.Data.Status,
		Date:                utility.RemoveZeroWidth(req.Data.Date),
		Location:            req.Data.Location,
		IsOtherSwabLocation: req.Data.IsOtherSwabLocation,
	}
	err := s.processAndValidateReq(swab)
	if err != nil {
		return nil, err
	}

	rslt, err := s.Model.CreateSwab(ctx, swab, constants.UserPatientMap[user.Role], user)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.SwabAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.SwabToResponse(rslt)
	return resp, nil
}

func (s *CreateSwabHandler) processAndValidateReq(swab *dto.Swab) error {
	swab.PatientID = utility.NormalizeID(swab.PatientID)
	if swab.PatientID == "" {
		return constants.InvalidPatientIDError
	}
	var err error
	swab.Date, err = utility.NormalizeDate(swab.Date)
	if err != nil {
		return err
	}
	if swab.Status < 1 || swab.Status > 3 {
		return constants.InvalidSwabStatusError
	}
	return nil
}
