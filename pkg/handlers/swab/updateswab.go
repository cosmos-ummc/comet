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

type UpdateSwabHandler struct {
	Model model.IModel
}

func (s *UpdateSwabHandler) UpdateSwab(ctx context.Context, req *pb.CommonSwabRequest, user *dto.User) (*pb.CommonSwabResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	swab := s.reqToSwab(req)

	err := s.processAndValidateReq(swab)
	if err != nil {
		return nil, err
	}

	v, err := s.Model.UpdateSwab(ctx, swab, constants.UserPatientMap[user.Role], user)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.SwabToResponse(v)
	return resp, nil
}

func (s *UpdateSwabHandler) reqToSwab(req *pb.CommonSwabRequest) *dto.Swab {
	swab := &dto.Swab{
		ID:                  utility.RemoveZeroWidth(req.Id),
		PatientID:           utility.RemoveZeroWidth(req.Data.PatientId),
		Status:              req.Data.Status,
		Location:            req.Data.Location,
		IsOtherSwabLocation: req.Data.IsOtherSwabLocation,
	}
	return swab
}

func (s *UpdateSwabHandler) processAndValidateReq(swab *dto.Swab) error {
	swab.PatientID = utility.NormalizeID(swab.PatientID)
	patientID, _ := utility.GetPatientIDAndDateFromSwabID(swab.ID)
	if swab.PatientID != patientID {
		return constants.InvalidPatientIDError
	}
	if swab.Status < 1 || swab.Status > 3 {
		return constants.InvalidSwabStatusError
	}
	return nil
}
