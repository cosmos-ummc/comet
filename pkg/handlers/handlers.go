package handlers

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	chatMessage "comet/pkg/handlers/chatmessage"
	chatRoom "comet/pkg/handlers/chatroom"
	"comet/pkg/handlers/consultant"
	"comet/pkg/handlers/declaration"
	"comet/pkg/handlers/feed"
	"comet/pkg/handlers/game"
	"comet/pkg/handlers/meditation"
	"comet/pkg/handlers/meeting"
	"comet/pkg/handlers/patient"
	"comet/pkg/handlers/question"
	"comet/pkg/handlers/report"
	"comet/pkg/handlers/tip"
	"comet/pkg/handlers/user"
	"comet/pkg/logger"
	"comet/pkg/model"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"os"
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
	resp, err := handler.CreatePatient(ctx, req)
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
	resp, err := handler.GetPatient(ctx, req)
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
	resp, err := handler.GetPatients(ctx, req)
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
	resp, err := handler.GetUndeclaredPatients(ctx, req)
	if err != nil {
		logger.Log.Error("GetUndeclaredPatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetUndeclaredPatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetStablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetStablePatientsHandler{Model: s.Model}
	resp, err := handler.GetStablePatients(ctx, req)
	if err != nil {
		logger.Log.Error("GetStablePatientsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetStablePatientsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetUnstablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &patient.GetStablePatientsHandler{Model: s.Model}
	resp, err := handler.GetStablePatients(ctx, req)
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
	resp, err := handler.UpdatePatient(ctx, req)
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
	resp, err := handler.UpdatePatients(ctx, req)
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
	resp, err := handler.DeletePatient(ctx, req)
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
	resp, err := handler.DeletePatients(ctx, req)
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

func (s *Handlers) ClientSetVisibility(ctx context.Context, req *pb.ClientSetVisibilityRequest) (*pb.ClientSetVisibilityResponse, error) {
	handler := &user.ClientSetVisibilityHandler{Model: s.Model}
	resp, err := handler.ClientSetVisibility(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientNewMatch(ctx context.Context, req *pb.ClientNewMatchRequest) (*pb.ClientNewMatchResponse, error) {
	handler := &user.ClientNewMatchHandler{Model: s.Model}
	resp, err := handler.ClientNewMatch(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientVerifyPatientComplete(ctx context.Context, req *pb.ClientVerifyPatientCompleteRequest) (*pb.ClientVerifyPatientCompleteResponse, error) {
	handler := &patient.ClientVerifyPatientCompleteHandler{Model: s.Model}
	resp, err := handler.ClientVerifyPatientComplete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientMessageEvent(ctx context.Context, req *pb.ClientMessageEventRequest) (*pb.ClientMessageEventResponse, error) {
	handler := &patient.ClientMessageEventHandler{Model: s.Model}
	resp, err := handler.ClientMessageEvent(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientGetChatRooms(ctx context.Context, req *pb.ClientGetChatRoomsRequest) (*pb.CommonChatRoomsResponse, error) {
	handler := &chatRoom.ClientGetChatRoomsHandler{Model: s.Model}
	resp, err := handler.ClientGetChatRooms(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientBlock(ctx context.Context, req *pb.ClientBlockRequest) (*pb.ClientBlockResponse, error) {
	handler := &user.ClientBlockHandler{Model: s.Model}
	resp, err := handler.ClientBlock(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientCheckCompleted(ctx context.Context, req *pb.ClientCheckCompletedRequest) (*pb.ClientCheckCompletedResponse, error) {
	handler := &user.ClientCheckCompletedHandler{Model: s.Model}
	resp, err := handler.ClientCheckCompleted(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) ClientSetNotFirstTime(ctx context.Context, req *pb.ClientSetNotFirstTimeRequest) (*pb.ClientSetNotFirstTimeResponse, error) {
	handler := &user.ClientSetNotFirstTimeHandler{Model: s.Model}
	resp, err := handler.ClientSetNotFirstTime(ctx, req)
	if err != nil {
		return nil, err
	}
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

func (s *Handlers) CreateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &question.CreateQuestionHandler{Model: s.Model}
	resp, err := handler.CreateQuestion(ctx, req)
	if err != nil {
		logger.Log.Error("CreateQuestionHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateQuestionHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetQuestion(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonQuestionResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.GetQuestionHandler{Model: s.Model}
	resp, err := handler.GetQuestion(ctx, req)
	if err != nil {
		logger.Log.Error("GetQuestionHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetQuestionHandler", zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
	return resp, nil
}

func (s *Handlers) GetQuestions(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonQuestionsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.GetQuestionsHandler{Model: s.Model}
	resp, err := handler.GetQuestions(ctx, req)
	if err != nil {
		logger.Log.Error("GetQuestionsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetQuestionsHandler", zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.UpdateQuestionHandler{Model: s.Model}
	resp, err := handler.UpdateQuestion(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateQuestionHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateQuestionHandler", zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateQuestions(ctx context.Context, req *pb.CommonQuestionsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.UpdateQuestionsHandler{Model: s.Model}
	resp, err := handler.UpdateQuestions(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateQuestionsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateQuestionsHandler", zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteQuestion(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonQuestionResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.DeleteQuestionHandler{Model: s.Model}
	resp, err := handler.DeleteQuestion(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteQuestionHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteQuestionHandler", zap.String("UserID", u.ID), zap.String("QuestionID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteQuestions(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &question.DeleteQuestionsHandler{Model: s.Model}
	resp, err := handler.DeleteQuestions(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteQuestionsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteQuestionsHandler", zap.String("UserID", u.ID), zap.Strings("QuestionIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &feed.CreateFeedHandler{Model: s.Model}
	resp, err := handler.CreateFeed(ctx, req)
	if err != nil {
		logger.Log.Error("CreateFeedHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateFeedHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetFeed(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonFeedResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.GetFeedHandler{Model: s.Model}
	resp, err := handler.GetFeed(ctx, req)
	if err != nil {
		logger.Log.Error("GetFeedHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetFeedHandler", zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
	return resp, nil
}

func (s *Handlers) GetFeeds(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonFeedsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.GetFeedsHandler{Model: s.Model}
	resp, err := handler.GetFeeds(ctx, req)
	if err != nil {
		logger.Log.Error("GetFeedsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetFeedsHandler", zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) ClientGetFeeds(ctx context.Context, req *pb.ClientGetFeedsRequest) (*pb.CommonFeedsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.ClientGetFeedsHandler{Model: s.Model}
	resp, err := handler.GetFeeds(ctx, req)
	if err != nil {
		logger.Log.Error("GetFeedsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetFeedsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.UpdateFeedHandler{Model: s.Model}
	resp, err := handler.UpdateFeed(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateFeedHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateFeedHandler", zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateFeeds(ctx context.Context, req *pb.CommonFeedsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.UpdateFeedsHandler{Model: s.Model}
	resp, err := handler.UpdateFeeds(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateFeedsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateFeedsHandler", zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteFeed(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonFeedResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.DeleteFeedHandler{Model: s.Model}
	resp, err := handler.DeleteFeed(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteFeedHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteFeedHandler", zap.String("UserID", u.ID), zap.String("FeedID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteFeeds(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &feed.DeleteFeedsHandler{Model: s.Model}
	resp, err := handler.DeleteFeeds(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteFeedsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteFeedsHandler", zap.String("UserID", u.ID), zap.Strings("FeedIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &tip.CreateTipHandler{Model: s.Model}
	resp, err := handler.CreateTip(ctx, req)
	if err != nil {
		logger.Log.Error("CreateTipHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateTipHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetTip(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonTipResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.GetTipHandler{Model: s.Model}
	resp, err := handler.GetTip(ctx, req)
	if err != nil {
		logger.Log.Error("GetTipHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("TipID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetTipHandler", zap.String("UserID", u.ID), zap.String("TipID", req.Id))
	return resp, nil
}

func (s *Handlers) GetTips(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonTipsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.GetTipsHandler{Model: s.Model}
	resp, err := handler.GetTips(ctx, req)
	if err != nil {
		logger.Log.Error("GetTipsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetTipsHandler", zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) ClientGetTips(ctx context.Context, req *pb.ClientGetTipsRequest) (*pb.CommonTipsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.ClientGetTipsHandler{Model: s.Model}
	resp, err := handler.GetTips(ctx, req)
	if err != nil {
		logger.Log.Error("GetTipsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetTipsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.UpdateTipHandler{Model: s.Model}
	resp, err := handler.UpdateTip(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateTipHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("TipID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateTipHandler", zap.String("UserID", u.ID), zap.String("TipID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateTips(ctx context.Context, req *pb.CommonTipsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.UpdateTipsHandler{Model: s.Model}
	resp, err := handler.UpdateTips(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateTipsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateTipsHandler", zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteTip(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonTipResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.DeleteTipHandler{Model: s.Model}
	resp, err := handler.DeleteTip(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteTipHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("TipID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteTipHandler", zap.String("UserID", u.ID), zap.String("TipID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteTips(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &tip.DeleteTipsHandler{Model: s.Model}
	resp, err := handler.DeleteTips(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteTipsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteTipsHandler", zap.String("UserID", u.ID), zap.Strings("TipIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &game.CreateGameHandler{Model: s.Model}
	resp, err := handler.CreateGame(ctx, req)
	if err != nil {
		logger.Log.Error("CreateGameHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateGameHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetGame(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonGameResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.GetGameHandler{Model: s.Model}
	resp, err := handler.GetGame(ctx, req)
	if err != nil {
		logger.Log.Error("GetGameHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("GameID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetGameHandler", zap.String("UserID", u.ID), zap.String("GameID", req.Id))
	return resp, nil
}

func (s *Handlers) GetGames(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonGamesResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.GetGamesHandler{Model: s.Model}
	resp, err := handler.GetGames(ctx, req)
	if err != nil {
		logger.Log.Error("GetGamesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetGamesHandler", zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) ClientGetGames(ctx context.Context, req *pb.ClientGetGamesRequest) (*pb.CommonGamesResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.ClientGetGamesHandler{Model: s.Model}
	resp, err := handler.GetGames(ctx, req)
	if err != nil {
		logger.Log.Error("GetGamesHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetGamesHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.UpdateGameHandler{Model: s.Model}
	resp, err := handler.UpdateGame(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateGameHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("GameID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateGameHandler", zap.String("UserID", u.ID), zap.String("GameID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateGames(ctx context.Context, req *pb.CommonGamesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.UpdateGamesHandler{Model: s.Model}
	resp, err := handler.UpdateGames(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateGamesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateGamesHandler", zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteGame(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonGameResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.DeleteGameHandler{Model: s.Model}
	resp, err := handler.DeleteGame(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteGameHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("GameID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteGameHandler", zap.String("UserID", u.ID), zap.String("GameID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteGames(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &game.DeleteGamesHandler{Model: s.Model}
	resp, err := handler.DeleteGames(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteGamesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteGamesHandler", zap.String("UserID", u.ID), zap.Strings("GameIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &meditation.CreateMeditationHandler{Model: s.Model}
	resp, err := handler.CreateMeditation(ctx, req)
	if err != nil {
		logger.Log.Error("CreateMeditationHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateMeditationHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetMeditation(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeditationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.GetMeditationHandler{Model: s.Model}
	resp, err := handler.GetMeditation(ctx, req)
	if err != nil {
		logger.Log.Error("GetMeditationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetMeditationHandler", zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
	return resp, nil
}

func (s *Handlers) GetMeditations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonMeditationsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.GetMeditationsHandler{Model: s.Model}
	resp, err := handler.GetMeditations(ctx, req)
	if err != nil {
		logger.Log.Error("GetMeditationsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetMeditationsHandler", zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) ClientGetMeditations(ctx context.Context, req *pb.ClientGetMeditationsRequest) (*pb.CommonMeditationsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.ClientGetMeditationsHandler{Model: s.Model}
	resp, err := handler.GetMeditations(ctx, req)
	if err != nil {
		logger.Log.Error("GetMeditationsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetMeditationsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) UpdateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.UpdateMeditationHandler{Model: s.Model}
	resp, err := handler.UpdateMeditation(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateMeditationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateMeditationHandler", zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateMeditations(ctx context.Context, req *pb.CommonMeditationsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.UpdateMeditationsHandler{Model: s.Model}
	resp, err := handler.UpdateMeditations(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateMeditationsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateMeditationsHandler", zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteMeditation(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeditationResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.DeleteMeditationHandler{Model: s.Model}
	resp, err := handler.DeleteMeditation(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteMeditationHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteMeditationHandler", zap.String("UserID", u.ID), zap.String("MeditationID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteMeditations(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meditation.DeleteMeditationsHandler{Model: s.Model}
	resp, err := handler.DeleteMeditations(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteMeditationsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteMeditationsHandler", zap.String("UserID", u.ID), zap.Strings("MeditationIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &chatMessage.CreateChatMessageHandler{Model: s.Model}
	resp, err := handler.CreateChatMessage(ctx, req)
	if err != nil {
		logger.Log.Error("CreateChatMessageHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateChatMessageHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetChatMessage(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatMessageResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.GetChatMessageHandler{Model: s.Model}
	resp, err := handler.GetChatMessage(ctx, req)
	if err != nil {
		logger.Log.Error("GetChatMessageHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetChatMessageHandler", zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
	return resp, nil
}

func (s *Handlers) GetChatMessages(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonChatMessagesResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.GetChatMessagesHandler{Model: s.Model}
	resp, err := handler.GetChatMessages(ctx, req)
	if err != nil {
		logger.Log.Error("GetChatMessagesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetChatMessagesHandler", zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.UpdateChatMessageHandler{Model: s.Model}
	resp, err := handler.UpdateChatMessage(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateChatMessageHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateChatMessageHandler", zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateChatMessages(ctx context.Context, req *pb.CommonChatMessagesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.UpdateChatMessagesHandler{Model: s.Model}
	resp, err := handler.UpdateChatMessages(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateChatMessagesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateChatMessagesHandler", zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteChatMessage(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatMessageResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.DeleteChatMessageHandler{Model: s.Model}
	resp, err := handler.DeleteChatMessage(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteChatMessageHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteChatMessageHandler", zap.String("UserID", u.ID), zap.String("ChatMessageID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteChatMessages(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatMessage.DeleteChatMessagesHandler{Model: s.Model}
	resp, err := handler.DeleteChatMessages(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteChatMessagesHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteChatMessagesHandler", zap.String("UserID", u.ID), zap.Strings("ChatMessageIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &chatRoom.CreateChatRoomHandler{Model: s.Model}
	resp, err := handler.CreateChatRoom(ctx, req)
	if err != nil {
		logger.Log.Error("CreateChatRoomHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateChatRoomHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetChatRoom(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatRoomResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.GetChatRoomHandler{Model: s.Model}
	resp, err := handler.GetChatRoom(ctx, req)
	if err != nil {
		logger.Log.Error("GetChatRoomHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetChatRoomHandler", zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
	return resp, nil
}

func (s *Handlers) GetChatRooms(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonChatRoomsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.GetChatRoomsHandler{Model: s.Model}
	resp, err := handler.GetChatRooms(ctx, req)
	if err != nil {
		logger.Log.Error("GetChatRoomsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetChatRoomsHandler", zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.UpdateChatRoomHandler{Model: s.Model}
	resp, err := handler.UpdateChatRoom(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateChatRoomHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateChatRoomHandler", zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateChatRooms(ctx context.Context, req *pb.CommonChatRoomsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.UpdateChatRoomsHandler{Model: s.Model}
	resp, err := handler.UpdateChatRooms(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateChatRoomsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateChatRoomsHandler", zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteChatRoom(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatRoomResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.DeleteChatRoomHandler{Model: s.Model}
	resp, err := handler.DeleteChatRoom(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteChatRoomHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteChatRoomHandler", zap.String("UserID", u.ID), zap.String("ChatRoomID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteChatRooms(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &chatRoom.DeleteChatRoomsHandler{Model: s.Model}
	resp, err := handler.DeleteChatRooms(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteChatRoomsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteChatRoomsHandler", zap.String("UserID", u.ID), zap.Strings("ChatRoomIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &consultant.CreateConsultantHandler{Model: s.Model}
	resp, err := handler.CreateConsultant(ctx, req)
	if err != nil {
		logger.Log.Error("CreateConsultantHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateConsultantHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetConsultant(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonConsultantResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.GetConsultantHandler{Model: s.Model}
	resp, err := handler.GetConsultant(ctx, req)
	if err != nil {
		logger.Log.Error("GetConsultantHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetConsultantHandler", zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
	return resp, nil
}

func (s *Handlers) GetConsultants(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonConsultantsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.GetConsultantsHandler{Model: s.Model}
	resp, err := handler.GetConsultants(ctx, req)
	if err != nil {
		logger.Log.Error("GetConsultantsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetConsultantsHandler", zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.UpdateConsultantHandler{Model: s.Model}
	resp, err := handler.UpdateConsultant(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateConsultantHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateConsultantHandler", zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateConsultants(ctx context.Context, req *pb.CommonConsultantsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.UpdateConsultantsHandler{Model: s.Model}
	resp, err := handler.UpdateConsultants(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateConsultantsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateConsultantsHandler", zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteConsultant(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonConsultantResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.DeleteConsultantHandler{Model: s.Model}
	resp, err := handler.DeleteConsultant(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteConsultantHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteConsultantHandler", zap.String("UserID", u.ID), zap.String("ConsultantID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteConsultants(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &consultant.DeleteConsultantsHandler{Model: s.Model}
	resp, err := handler.DeleteConsultants(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteConsultantsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteConsultantsHandler", zap.String("UserID", u.ID), zap.Strings("ConsultantIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) CreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &meeting.CreateMeetingHandler{Model: s.Model}
	resp, err := handler.CreateMeeting(ctx, req)
	if err != nil {
		logger.Log.Error("CreateMeetingHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("CreateMeetingHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) ClientCreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	handler := &meeting.ClientCreateMeetingHandler{Model: s.Model}
	resp, err := handler.ClientCreateMeeting(ctx, req)
	if err != nil {
		logger.Log.Error("CreateMeetingHandler: " + err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *Handlers) GetMeeting(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeetingResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.GetMeetingHandler{Model: s.Model}
	resp, err := handler.GetMeeting(ctx, req)
	if err != nil {
		logger.Log.Error("GetMeetingHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
		return nil, err
	}
	logger.Log.Info("GetMeetingHandler", zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
	return resp, nil
}

func (s *Handlers) GetMeetings(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonMeetingsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.GetMeetingsHandler{Model: s.Model}
	resp, err := handler.GetMeetings(ctx, req)
	if err != nil {
		logger.Log.Error("GetMeetingsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("GetMeetingsHandler", zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) UpdateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.UpdateMeetingHandler{Model: s.Model}
	resp, err := handler.UpdateMeeting(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateMeetingHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
		return nil, err
	}
	logger.Log.Info("UpdateMeetingHandler", zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
	return resp, nil
}

func (s *Handlers) UpdateMeetings(ctx context.Context, req *pb.CommonMeetingsRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.UpdateMeetingsHandler{Model: s.Model}
	resp, err := handler.UpdateMeetings(ctx, req)
	if err != nil {
		logger.Log.Error("UpdateMeetingsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("UpdateMeetingsHandler", zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
	return resp, nil
}

func (s *Handlers) DeleteMeeting(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeetingResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.DeleteMeetingHandler{Model: s.Model}
	resp, err := handler.DeleteMeeting(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteMeetingHandler: "+err.Error(), zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
		return nil, err
	}
	logger.Log.Info("DeleteMeetingHandler", zap.String("UserID", u.ID), zap.String("MeetingID", req.Id))
	return resp, nil
}

func (s *Handlers) DeleteMeetings(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	u, err := s.validateUser(ctx, constants.SuperUserOnly)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &meeting.DeleteMeetingsHandler{Model: s.Model}
	resp, err := handler.DeleteMeetings(ctx, req)
	if err != nil {
		logger.Log.Error("DeleteMeetingsHandler: "+err.Error(), zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
		return nil, err
	}
	logger.Log.Info("DeleteMeetingsHandler", zap.String("UserID", u.ID), zap.Strings("MeetingIDs", req.Ids))
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
	resp, err := handler.CreateDeclaration(ctx, req)
	if err != nil {
		logger.Log.Error("CreateDeclarationHandler: "+err.Error(),
			zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId))
		return nil, err
	}
	logger.Log.Info("CreateDeclarationHandler", zap.String("UserID", u.ID), zap.String("PatientID", req.Data.PatientId))
	return resp, nil
}

func (s *Handlers) GetDeclaration(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonDeclarationResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.GetDeclarationHandler{Model: s.Model}
	resp, err := handler.GetDeclaration(ctx, req)
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
	resp, err := handler.GetDeclarations(ctx, req)
	if err != nil {
		logger.Log.Error("GetDeclarationsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetDeclarationsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetNormalDeclarations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.GetNormalDeclarationsHandler{Model: s.Model}
	resp, err := handler.GetNormalDeclarations(ctx, req)
	if err != nil {
		logger.Log.Error("GetDeclarationsHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetDeclarationsHandler", zap.String("UserID", u.ID))
	return resp, nil
}

func (s *Handlers) GetDailyDeclarations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error) {
	u, err := s.validateUser(ctx, constants.AllCanAccess)
	if err != nil {
		return nil, constants.UnauthorizedAccessError
	}
	handler := &declaration.GetDailyDeclarationsHandler{Model: s.Model}
	resp, err := handler.GetDailyDeclarations(ctx, req)
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
	resp, err := handler.UpdateDeclaration(ctx, req)
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
	resp, err := handler.UpdateDeclarations(ctx, req)
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
	resp, err := handler.DeleteDeclaration(ctx, req)
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
	resp, err := handler.DeleteDeclarations(ctx, req)
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
	resp, err := handler.GetReport(ctx, req)
	if err != nil {
		logger.Log.Error("GetReportHandler: "+err.Error(), zap.String("UserID", u.ID))
		return nil, err
	}
	logger.Log.Info("GetReportHandler", zap.String("UserID", u.ID))
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

func (s *Handlers) validateUser(ctx context.Context, roles []string) (*dto.User, error) {
	if os.Getenv("AUTH_ENABLED") != "true" {
		return &dto.User{}, nil
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

	// get user
	u, err := s.Model.GetUser(ctx, token)
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
