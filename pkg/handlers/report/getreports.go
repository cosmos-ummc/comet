package report

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetReportsHandler struct {
	Model model.IModel
}

func (s *GetReportsHandler) GetReports(ctx context.Context, req *pb.GetReportsRequest) (*pb.CommonReportsResponse, error) {
	s.processReq(req)
	reports, err := s.Model.GetReports(ctx, req.From, req.To)
	if err != nil {
		return nil, err
	}
	resp := utility.ReportsToResponse(reports)
	return resp, nil
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
