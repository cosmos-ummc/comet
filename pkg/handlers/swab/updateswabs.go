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

type UpdateSwabsHandler struct {
	Model model.IModel
}

func (s *UpdateSwabsHandler) UpdateSwabs(ctx context.Context, req *pb.CommonSwabsRequest, user *dto.User) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	swab := s.reqToSwab(req)

	err := s.processAndValidateReq(swab, req.Ids)
	if err != nil {
		return nil, err
	}

	ids, err := s.Model.UpdateSwabs(ctx, swab, req.Ids, constants.UserPatientMap[user.Role], user)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateSwabsHandler) reqToSwab(req *pb.CommonSwabsRequest) *dto.Swab {
	swab := &dto.Swab{
		PatientID:           utility.RemoveZeroWidth(req.Data.PatientId),
		Status:              req.Data.Status,
		Location:            req.Data.Location,
		IsOtherSwabLocation: req.Data.IsOtherSwabLocation,
	}
	return swab
}

func (s *UpdateSwabsHandler) processAndValidateReq(swab *dto.Swab, ids []string) error {
	swab.PatientID = utility.NormalizeID(swab.PatientID)
	if swab.Status < 1 || swab.Status > 3 {
		return constants.InvalidSwabStatusError
	}
	for _, id := range ids {
		patientID, _ := utility.GetPatientIDAndDateFromSwabID(id)
		if swab.PatientID != patientID {
			return constants.InvalidPatientIDError
		}
	}
	return nil
}
