package handlers

import (
	pb "comet/pkg/api"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// IHandlers ...
type IHandlers interface {
	CreatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error)
	GetPatient(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonPatientResponse, error)
	GetPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error)
	GetUndeclaredPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error)
	GetUnstablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error)
	GetStablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error)
	UpdatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error)
	UpdatePatients(ctx context.Context, req *pb.CommonPatientsRequest) (*pb.CommonIdsResponse, error)
	DeletePatient(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonPatientResponse, error)
	DeletePatients(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error)
	GetUser(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonUserResponse, error)
	GetUsers(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonUsersResponse, error)
	UpdateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error)
	UpdateUsers(ctx context.Context, req *pb.CommonUsersRequest) (*pb.CommonIdsResponse, error)
	DeleteUser(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonUserResponse, error)
	DeleteUsers(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error)
	GetDeclaration(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonDeclarationResponse, error)
	GetDeclarations(ctx context.Context, in *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error)
	UpdateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error)
	UpdateDeclarations(ctx context.Context, req *pb.CommonDeclarationsRequest) (*pb.CommonIdsResponse, error)
	DeleteDeclaration(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonDeclarationResponse, error)
	DeleteDeclarations(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.CommonReportResponse, error)
	GetReports(ctx context.Context, req *pb.GetReportsRequest) (*pb.CommonReportsResponse, error)

	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, req *empty.Empty) (*empty.Empty, error)
	Refresh(ctx context.Context, req *empty.Empty) (*pb.RefreshResponse, error)
	UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*empty.Empty, error)
	GetPasswordReset(ctx context.Context, req *pb.GetPasswordResetRequest) (*pb.GetPasswordResetResponse, error)

	ClientGetPatients(ctx context.Context, req *pb.ClientGetPatientsRequest) (*pb.ClientGetPatientsResponse, error)
	ClientGetPatientV2(ctx context.Context, req *pb.ClientGetPatientV2Request) (*pb.ClientGetPatientV2Response, error)
	ClientGetUndeclaredPatients(ctx context.Context, req *pb.ClientGetUndeclaredPatientsRequest) (*pb.ClientGetUndeclaredPatientsResponse, error)
	ClientUpdatePatient(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*empty.Empty, error)
	ClientUpdatePatientV2(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*pb.ClientUpdatePatientV2Response, error)
	ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error)

	CreateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error)
	GetQuestion(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonQuestionResponse, error)
	GetQuestions(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error)
	UpdateQuestions(ctx context.Context, req *pb.CommonQuestionsRequest) (*pb.CommonIdsResponse, error)
	DeleteQuestion(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonQuestionResponse, error)
	DeleteQuestions(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error)
	GetChatMessage(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatMessageResponse, error)
	GetChatMessages(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonChatMessagesResponse, error)
	UpdateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error)
	UpdateChatMessages(ctx context.Context, req *pb.CommonChatMessagesRequest) (*pb.CommonIdsResponse, error)
	DeleteChatMessage(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatMessageResponse, error)
	DeleteChatMessages(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error)
	GetChatRoom(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatRoomResponse, error)
	GetChatRooms(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonChatRoomsResponse, error)
	UpdateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error)
	UpdateChatRooms(ctx context.Context, req *pb.CommonChatRoomsRequest) (*pb.CommonIdsResponse, error)
	DeleteChatRoom(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatRoomResponse, error)
	DeleteChatRooms(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)
}
