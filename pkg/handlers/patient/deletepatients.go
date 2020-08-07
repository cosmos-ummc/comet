package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type DeletePatientsHandler struct {
	Model model.IModel
}

func (s *DeletePatientsHandler) DeletePatients(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	req.Ids = s.processReq(req.Ids)

	deletedIDs, err := s.Model.DeletePatients(ctx, req.Ids)
	if err != nil{
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: deletedIDs}, nil
}

func (s *DeletePatientsHandler) processReq(ids []string) []string {
	split := strings.Split(ids[0], ",")
	var newIDs []string
	for _, id := range split {
		newIDs = append(newIDs, utility.NormalizeID(id))
	}
	return newIDs
}
