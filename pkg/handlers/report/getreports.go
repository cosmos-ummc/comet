package report

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetReportsHandler struct {
	Model model.IModel
}

func (s *GetReportsHandler) GetReports(ctx context.Context, req *pb.GetReportsRequest, user *dto.User) (*pb.CommonReportsResponse, error) {
	s.processReq(req)

	if req.Type == constants.CallingReportCode {
		reports, err := s.Model.GetCallingReports(ctx, req.From, req.To, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.CallingReportsToResponse(reports)
		return resp, nil
	}
	if req.Type == constants.DeclarationReportCode {
		reports, err := s.Model.GetDeclarationReports(ctx, req.From, req.To, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.DeclarationReportsToResponse(reports)
		return resp, nil
	}
	if req.Type == constants.PatientStatusReportCode {
		reports, err := s.Model.GetPatientStatusReports(ctx, req.From, req.To, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, err
		}
		resp := utility.PatientStatusReportsToResponse(reports)
		return resp, nil
	}

	return &pb.CommonReportsResponse{}, nil
}

func (s *GetReportsHandler) processReq(req *pb.GetReportsRequest) {
	var err error
	req.From, err = utility.NormalizeDate(req.From)
	if err != nil {
		return
	}
	req.To, err = utility.NormalizeDate(req.To)
	if err != nil {
		return
	}
}
