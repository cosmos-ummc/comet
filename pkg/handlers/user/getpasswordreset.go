package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetPasswordResetHandler struct {
	Model model.IModel
}

func (s *GetPasswordResetHandler) GetPasswordReset(ctx context.Context, req *pb.GetPasswordResetRequest) (*pb.GetPasswordResetResponse, error) {

	// get user
	user, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.UserNotFoundError
		}
		return nil, constants.InternalError
	}

	auth, err := s.Model.CreateToken(ctx, &dto.AuthObject{
		UserId:      req.Id,
		TTL:         utility.MilliToTime(time.Now().Add(time.Hour*24*constants.PasswordResetTokenTTLDays).Unix()*1000 - 1000),
		DisplayName: user.Name,
		Type:        constants.Reset,
	})
	if err != nil {
		logger.Log.Error("GetPasswordReset: " + err.Error())
		return nil, err
	}

	passwordReset := os.Getenv("ADMIN_URL") + "/#/resetpassword?token=" + auth.Token

	// send password reset email
	err = utility.SendPasswordResetEmail(user.Email, user.Name, passwordReset)
	if err != nil {
		logger.Log.Warn("GetPasswordReset: " + err.Error())
		return nil, err
	}

	return &pb.GetPasswordResetResponse{
		Message: "Password reset link has been sent to the user's email.",
	}, nil
}
