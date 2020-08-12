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
	report, err := s.Model.GetReport(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := utility.ReportToResponse(report)
	return resp, nil
}
