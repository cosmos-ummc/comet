package swab

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteSwabsHandler struct {
	Model model.IModel
}

func (s *DeleteSwabsHandler) DeleteSwabs(ctx context.Context, req *pb.CommonDeletesRequest, user *dto.User) (*pb.CommonIdsResponse, error) {
	req.Ids = s.processReq(req.Ids)
	ids, err := s.Model.DeleteSwabs(ctx, req.Ids, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *DeleteSwabsHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
