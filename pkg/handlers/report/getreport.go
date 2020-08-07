package report

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetReportHandler struct {
	Model model.IModel
}

func (s *GetReportHandler) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.CommonReportResponse, error) {
	err := s.processReq(req)
	if err != nil {
		return nil, err
	}

	report, err := s.Model.GetReport(ctx, req.Date)
	if err != nil {
		return nil, err
	}

	resp := utility.ReportToResponse(report)
	return resp, nil
}

func (s *GetReportHandler) processReq(req *pb.GetReportRequest) error {
	var err error
	req.Date, err = utility.NormalizeDate(req.Date)
	return err
}
