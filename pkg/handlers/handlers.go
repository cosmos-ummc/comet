package handlers

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/handlers/activity"
	"comet/pkg/handlers/declaration"
	"comet/pkg/handlers/patient"
	"comet/pkg/handlers/report"
	"comet/pkg/handlers/swab"
	"comet/pkg/handlers/user"
	"comet/pkg/logger"
	"comet/pkg/model"
	"context"
	"errors"
	"os"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// Handlers ...
type Handlers struct {
	Model model.IModel
}

// NewHandlers ...
func NewHandlers(model model.IModel) IHandlers {
	return &Handlers{Model: model}
}

func (s *Handlers) CreatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.CreatePatientHandler{Model: s.Model}
	resp, err := handler.CreatePatient(ctx, req, u)
	if err != nil {
		logger.Log.Error("CreatePatientHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("CreatePatientHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) GetPatient(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonPatientResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetPatientHandler{Model: s.Model}
	resp, err := handler.GetPatient(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetPatientHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetPatientHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) GetPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetPatientsHandler{Model: s.Model}
	resp, err := handler.GetPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetUndeclaredPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetUndeclaredPatientsHandler{Model: s.Model}
	resp, err := handler.GetUndeclaredPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetUndeclaredPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetUndeclaredPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetCallPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetCallPatientsHandler{Model: s.Model}
	resp, err := handler.GetCallPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetCallPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetCallPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetNoCallPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetNoCallPatientsHandler{Model: s.Model}
	resp, err := handler.GetNoCallPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetNoCallPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetNoCallPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetSwabPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetSwabPatientsHandler{Model: s.Model}
	resp, err := handler.GetSwabPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetSwabPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetSwabPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetOtherPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetOtherPatientsHandler{Model: s.Model}
	resp, err := handler.GetOtherPatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetOtherPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetOtherPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetStablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetStablePatientsHandler{Model: s.Model}
	resp, err := handler.GetStablePatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetStablePatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetStablePatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.UpdatePatientHandler{Model: s.Model}
	resp, err := handler.UpdatePatient(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdatePatientHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdatePatientHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdatePatients(ctx context.Context, req *pb.CommonPatientsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.UpdatePatientsHandler{Model: s.Model}
	resp, err := handler.UpdatePatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdatePatientsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("PatientIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdatePatientsHandler", zap.String("UserID", u.ID), zap.Strings("PatientIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeletePatient(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonPatientResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.DeletePatientHandler{Model: s.Model}
	resp, err := handler.DeletePatient(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeletePatientHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeletePatientHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) DeletePatients(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.DeletePatientsHandler{Model: s.Model}
	resp, err := handler.DeletePatients(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeletePatientsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("PatientIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeletePatientsHandler", zap.String("UserID", u.ID), zap.Strings("PatientIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.CreateUserHandler{Model: s.Model}
	resp, err := handler.CreateUser(ctx, req)
	if err != nil {
		logger.Log.Error("CreateUserHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateUserHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetUser(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonUserResponse, error) {
	handler := &user.GetUserHandler{Model: s.Model}
	resp, err := handler.GetUser(ctx, req)
	if err != nil {
		logger.Log.Error("GetUserHandler: "+err.Error(), zap.String("UserID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetUserHandler", zap.String("UserID", req.Id))
	return resp, nil
}

func (s *Handlers) GetUsers(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonUsersResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.GetUsersHandler{Model: s.Model}
	resp, err := handler.GetUsers(ctx, req)
	if err != nil {
		logger.Log.Error("GetUsersHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetUsersHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.UpdateUserHandler{Model: s.Model}
	resp, err := handler.UpdateUser(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateUserHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("TargetUserID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateUserHandler", zap.String("UserID", u.ID), zap.String("TargetUserID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateUsers(ctx context.Context, req *pb.CommonUsersRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.UpdateUsersHandler{Model: s.Model}
	resp, err := handler.UpdateUsers(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateUsersHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("TargetUserIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateUsersHandler", zap.String("UserID", u.ID), zap.Strings("TargetUserIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteUser(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonUserResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.DeleteUserHandler{Model: s.Model}
	resp, err := handler.DeleteUser(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteUserHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("TargetUserID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteUserHandler", zap.String("UserID", u.ID), zap.String("TargetUserID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteUsers(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &user.DeleteUsersHandler{Model: s.Model}
	resp, err := handler.DeleteUsers(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteUsersHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("TargetUserIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteUsersHandler", zap.String("UserID", u.ID), zap.Strings("TargetUserIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateSwab(ctx context.Context, req *pb.CommonSwabRequest) (*pb.CommonSwabResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &swab.CreateSwabHandler{Model: s.Model}
	resp, err := handler.CreateSwab(ctx, req, u)
	if err != nil {
		logger.Log.Error("CreateSwabHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId))
		return nil, err
	}
	logger.Log.Info("CreateSwabHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId))
	return resp, nil
}

func (s *Handlers) GetSwab(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonSwabResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.GetSwabHandler{Model: s.Model}
	resp, err := handler.GetSwab(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetSwabHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetSwabHandler", zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
	return resp, nil
}

func (s *Handlers) GetSwabs(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonSwabsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.GetSwabsHandler{Model: s.Model}
	resp, err := handler.GetSwabs(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetSwabsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetSwabsHandler", zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateSwab(ctx context.Context, req *pb.CommonSwabRequest) (*pb.CommonSwabResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.UpdateSwabHandler{Model: s.Model}
	resp, err := handler.UpdateSwab(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdateSwabHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateSwabHandler", zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateSwabs(ctx context.Context, req *pb.CommonSwabsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.UpdateSwabsHandler{Model: s.Model}
	resp, err := handler.UpdateSwabs(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdateSwabsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateSwabsHandler", zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteSwab(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonSwabResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.DeleteSwabHandler{Model: s.Model}
	resp, err := handler.DeleteSwab(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeleteSwabHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteSwabHandler", zap.String("UserID", u.ID), zap.String("SwabID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteSwabs(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &swab.DeleteSwabsHandler{Model: s.Model}
	resp, err := handler.DeleteSwabs(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeleteSwabsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteSwabsHandler", zap.String("UserID", u.ID), zap.Strings("SwabIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error) {
	u, err := s.validateUser(ctx, []string{})
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &declaration.CreateDeclarationHandler{Model: s.Model}
	resp, err := handler.CreateDeclaration(ctx, req, u)
	if err != nil {
		logger.Log.Error("CreateDeclarationHandler: "+err.Error(),
			zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId),
			zap.String("Date", req.Data.Date))
		return nil, err
	}
	logger.Log.Info("CreateDeclarationHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId),
		zap.String("Date", req.Data.Date))
	return resp, nil
}

func (s *Handlers) GetDeclaration(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonDeclarationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.GetDeclarationHandler{Model: s.Model}
	resp, err := handler.GetDeclaration(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetDeclarationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetDeclarationHandler", zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
	return resp, nil
}

func (s *Handlers) GetDeclarations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.GetDeclarationsHandler{Model: s.Model}
	resp, err := handler.GetDeclarations(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetDeclarationsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetDeclarationsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.UpdateDeclarationHandler{Model: s.Model}
	resp, err := handler.UpdateDeclaration(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdateDeclarationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateDeclarationHandler", zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateDeclarations(ctx context.Context, req *pb.CommonDeclarationsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.UpdateDeclarationsHandler{Model: s.Model}
	resp, err := handler.UpdateDeclarations(ctx, req, u)
	if err != nil {
		logger.Log.Error("UpdateDeclarationsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("DeclarationIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateDeclarationsHandler", zap.String("UserID", u.ID), zap.Strings("DeclarationIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteDeclaration(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonDeclarationResponse, error) {
	u, err := s.validateUser(ctx, []string{})
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.DeleteDeclarationHandler{Model: s.Model}
	resp, err := handler.DeleteDeclaration(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeleteDeclarationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteDeclarationHandler", zap.String("UserID", u.ID), zap.String("DeclarationID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteDeclarations(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, []string{})
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.DeleteDeclarationsHandler{Model: s.Model}
	resp, err := handler.DeleteDeclarations(ctx, req, u)
	if err != nil {
		logger.Log.Error("DeleteDeclarationsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("DeclarationIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteDeclarationsHandler", zap.String("UserID", u.ID), zap.Strings("DeclarationIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.CommonReportResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &report.GetReportHandler{Model: s.Model}
	resp, err := handler.GetReport(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetReportHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetReportHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetReports(ctx context.Context, req *pb.GetReportsRequest) (*pb.CommonReportsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &report.GetReportsHandler{Model: s.Model}
	resp, err := handler.GetReports(ctx, req, u)
	if err != nil {
		logger.Log.Error("GetReportsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetReportsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	handler := &user.LoginHandler{Model: s.Model}
	resp, err := handler.Login(ctx, req)
	if err != nil {
		logger.Log.Error("LoginHandler: "+err.Error(), zap.String("email", req.Email))
		return nil, err
	}
	logger.Log.Info("LoginHandler", zap.String("email", req.Email))
	return resp, nil
}

func (s *Handlers) Logout(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	handler := &user.LogoutHandler{Model: s.Model}
	resp, err := handler.Logout(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) Refresh(ctx context.Context, _ *empty.Empty) (*pb.RefreshResponse, error) {
	handler := &user.RefreshHandler{Model: s.Model}
	resp, err := handler.Refresh(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*empty.Empty, error) {
	handler := &user.UpdatePasswordHandler{Model: s.Model}
	resp, err := handler.UpdatePassword(ctx, req)
	if err != nil {
		logger.Log.Error("UpdatePasswordHandler: " + err.Error())
		return nil, err
	}
	logger.Log.Info("UpdatePasswordHandler")
	return resp, nil
}

func (s *Handlers) GetPasswordReset(ctx context.Context, req *pb.GetPasswordResetRequest) (*pb.GetPasswordResetResponse, error) {
	//u, err := s.validateUser(ctx, constants.SuperUserOnly)
	//if err != nil {
	//	return nil, constants.UnauthorizedAccessError
	//}
	handler := &user.GetPasswordResetHandler{Model: s.Model}
	resp, err := handler.GetPasswordReset(ctx, req)
	if err != nil {
		logger.Log.Error("GetPasswordResetHandler: "+err.Error(), zap.String("TargetUserID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetPasswordResetHandler", zap.String("TargetUserID", req.Id))
	return resp, nil
}

func (s *Handlers) ClientGetPatients(ctx context.Context, req *pb.ClientGetPatientsRequest) (*pb.ClientGetPatientsResponse, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.ClientGetPatientsHandler{Model: s.Model}
	resp, err := handler.ClientGetPatients(ctx, req)
	if err != nil {
		logger.Log.Error("ClientGetPatientsHandler: "+err.Error(), zap.String("TelegramID", req.TelegramId))
		return nil, err
	}
	logger.Log.Info("ClientGetPatientsHandler", zap.String("TelegramID", req.TelegramId))
	return resp, nil
}

func (s *Handlers) ClientGetPatientV2(ctx context.Context, req *pb.ClientGetPatientV2Request) (*pb.ClientGetPatientV2Response, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.ClientGetPatientV2Handler{Model: s.Model}
	resp, err := handler.ClientGetPatientV2(ctx, req)
	if err != nil {
		logger.Log.Error("ClientGetPatientV2Handler: "+err.Error(), zap.String("TelegramID", req.TelegramId))
		return nil, err
	}
	logger.Log.Info("ClientGetPatientV2Handler", zap.String("TelegramID", req.TelegramId))
	return resp, nil
}

func (s *Handlers) ClientGetUndeclaredPatients(ctx context.Context, req *pb.ClientGetUndeclaredPatientsRequest) (*pb.ClientGetUndeclaredPatientsResponse, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.ClientGetUndeclaredPatientsHandler{Model: s.Model}
	resp, err := handler.ClientGetUndeclaredPatients(ctx, req)
	if err != nil {
		logger.Log.Error("ClientGetUndeclaredPatientsHandler: "+err.Error(), zap.Int64("From", req.From), zap.Int64("To", req.To))
		return nil, err
	}
	logger.Log.Info("ClientGetUndeclaredPatientsHandler", zap.Int64("From", req.From), zap.Int64("To", req.To))
	return resp, nil
}

func (s *Handlers) ClientUpdatePatient(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*empty.Empty, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.ClientUpdatePatientHandler{Model: s.Model}
	resp, err := handler.ClientUpdatePatient(ctx, req)
	if err != nil {
		logger.Log.Error("ClientUpdatePatientHandler: "+err.Error(), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("ClientUpdatePatientHandler", zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) ClientUpdatePatientV2(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*pb.ClientUpdatePatientV2Response, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.ClientUpdatePatientV2Handler{Model: s.Model}
	resp, err := handler.ClientUpdatePatientV2(ctx, req)
	if err != nil {
		logger.Log.Error("ClientUpdatePatientV2Handler: "+err.Error(), zap.String("PatientID", req.Id))
		return nil, err
	}
	logger.Log.Info("ClientUpdatePatientV2Handler", zap.String("PatientID", req.Id))
	return resp, nil
}

func (s *Handlers) ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error) {
	_, err := s.validateUser(ctx, constants.ChatBotOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.ClientCreateDeclarationHandler{Model: s.Model}
	resp, err := handler.ClientCreateDeclaration(ctx, req)
	if err != nil {
		logger.Log.Error("ClientCreateDeclarationHandler: "+err.Error(), zap.String("PatientID", req.PatientId))
		return nil, err
	}
	logger.Log.Info("ClientCreateDeclarationHandler", zap.String("PatientID", req.PatientId))
	return resp, nil
}

func (s *Handlers) GetActivities(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonActivitiesResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &activity.GetActivitiesHandler{Model: s.Model}
	resp, err := handler.GetActivities(ctx, req)
	if err != nil {
		logger.Log.Error("GetActivitiesHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetActivitiesHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetActivity(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonActivityResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &activity.GetActivityHandler{Model: s.Model}
	resp, err := handler.GetActivity(ctx, req)
	if err != nil {
		logger.Log.Error("GetActivityHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ActivityID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetActivityHandler", zap.String("UserID", u.ID), zap.String("ActivityID", req.Id))
	return resp, nil
}

func (s *Handlers) validateUser(ctx context.Context, roles []string) (*dto.User, error) {
	if os.Getenv("AUTH_ENABLED") != "true" {
		return nil, nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("ValidateUser: metadata not found")
	}
	tokenSlice := md.Get("authorization")
	if len(tokenSlice) < 1 {
		return nil, errors.New("ValidateUser: token not found")
	}
	token := tokenSlice[0]

	// exemption: chat bot user
	if token == os.Getenv("CHATBOT_USER") {
		// check if user is allowed to access the API
		for _, role := range roles {
			if role == constants.ChatBot {
				return &dto.User{ID: token, Role: constants.ChatBot}, nil
			}
		}
		return nil, errors.New("unauthorized access")
	}

	// exemption: backend user
	if token == os.Getenv("BACKEND_USER") {
		return &dto.User{ID: token, Role: constants.Superuser}, nil
	}

	u, err := s.Model.VerifyUser(ctx, strings.Join(tokenSlice, " "))
	if err != nil {
		return nil, err
	}

	// check if user is allowed to access the API
	authorized := false
	for _, role := range roles {
		if u.Role == role {
			authorized = true
		}
	}

	if !authorized {
		return nil, errors.New("unauthorized access")
	}

	return u, nil
}
