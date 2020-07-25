package report

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetReportHandler struct {
	Model model.IModel
}

func (s *GetReportHandler) GetReport(ctx context.Context, req *pb.GetReportRequest, user *dto.User) (*pb.CommonReportResponse, error) {
	err := s.processReq(req)
	if err != nil {
		return nil, err
	}

	if req.Type == constants.CallingReportCode {
		report, err := s.Model.GetCallingReport(ctx, req.Date, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.CallingReportToResponse(report)
		return resp, nil
	}
	if req.Type == constants.DeclarationReportCode {
		report, err := s.Model.GetDeclarationReport(ctx, req.Date, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.DeclarationReportToResponse(report)
		return resp, nil
	}
	if req.Type == constants.PatientStatusReportCode {
		report, err := s.Model.GetPatientStatusReport(ctx, req.Date, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.PatientStatusReportToResponse(report)
		return resp, nil
	}

	return &pb.CommonReportResponse{}, nil
}

func (s *GetReportHandler) processReq(req *pb.GetReportRequest) error {
	var err error
	req.Date, err = utility.NormalizeDate(req.Date)
	return err
}
