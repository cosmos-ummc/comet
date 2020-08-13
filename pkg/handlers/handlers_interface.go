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
	ClientBlock(ctx context.Context, req *pb.ClientBlockRequest) (*pb.ClientBlockResponse, error)

	CreateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error)
	GetDeclaration(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonDeclarationResponse, error)
	GetDeclarations(ctx context.Context, in *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error)
	UpdateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error)
	UpdateDeclarations(ctx context.Context, req *pb.CommonDeclarationsRequest) (*pb.CommonIdsResponse, error)
	DeleteDeclaration(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonDeclarationResponse, error)
	DeleteDeclarations(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.CommonReportResponse, error)

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
	ClientNewMatch(ctx context.Context, req *pb.ClientNewMatchRequest) (*pb.ClientNewMatchResponse, error)
	ClientGetChatRooms(ctx context.Context, req *pb.ClientGetChatRoomsRequest) (*pb.CommonChatRoomsResponse, error)
	ClientCheckCompleted(ctx context.Context, req *pb.ClientCheckCompletedRequest) (*pb.ClientCheckCompletedResponse, error)
	ClientSetNotFirstTime(ctx context.Context, req *pb.ClientSetNotFirstTimeRequest) (*pb.ClientSetNotFirstTimeResponse, error)
	ClientCreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error)

	ClientGetTips(ctx context.Context, req *pb.ClientGetTipsRequest) (*pb.CommonTipsResponse, error)
	ClientGetFeeds(ctx context.Context, req *pb.ClientGetFeedsRequest) (*pb.CommonFeedsResponse, error)
	ClientGetGames(ctx context.Context, req *pb.ClientGetGamesRequest) (*pb.CommonGamesResponse, error)
	ClientGetMeditations(ctx context.Context, req *pb.ClientGetMeditationsRequest) (*pb.CommonMeditationsResponse, error)
	ClientSetVisibility(ctx context.Context, req *pb.ClientSetVisibilityRequest) (*pb.ClientSetVisibilityResponse, error)
	ClientVerifyPatientComplete(ctx context.Context, req *pb.ClientVerifyPatientCompleteRequest) (*pb.ClientVerifyPatientCompleteResponse, error)
	ClientMessageEvent(ctx context.Context, req *pb.ClientMessageEventRequest) (*pb.ClientMessageEventResponse, error)

	CreateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error)
	GetQuestion(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonQuestionResponse, error)
	GetQuestions(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, req *pb.CommonQuestionRequest) (*pb.CommonQuestionResponse, error)
	UpdateQuestions(ctx context.Context, req *pb.CommonQuestionsRequest) (*pb.CommonIdsResponse, error)
	DeleteQuestion(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonQuestionResponse, error)
	DeleteQuestions(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error)
	GetFeed(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonFeedResponse, error)
	GetFeeds(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonFeedsResponse, error)
	UpdateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error)
	UpdateFeeds(ctx context.Context, req *pb.CommonFeedsRequest) (*pb.CommonIdsResponse, error)
	DeleteFeed(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonFeedResponse, error)
	DeleteFeeds(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error)
	GetTip(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonTipResponse, error)
	GetTips(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonTipsResponse, error)
	UpdateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error)
	UpdateTips(ctx context.Context, req *pb.CommonTipsRequest) (*pb.CommonIdsResponse, error)
	DeleteTip(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonTipResponse, error)
	DeleteTips(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

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

	CreateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error)
	GetConsultant(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonConsultantResponse, error)
	GetConsultants(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonConsultantsResponse, error)
	UpdateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error)
	UpdateConsultants(ctx context.Context, req *pb.CommonConsultantsRequest) (*pb.CommonIdsResponse, error)
	DeleteConsultant(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonConsultantResponse, error)
	DeleteConsultants(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error)
	GetMeeting(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeetingResponse, error)
	GetMeetings(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonMeetingsResponse, error)
	UpdateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error)
	UpdateMeetings(ctx context.Context, req *pb.CommonMeetingsRequest) (*pb.CommonIdsResponse, error)
	DeleteMeeting(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeetingResponse, error)
	DeleteMeetings(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error)
	GetGame(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonGameResponse, error)
	GetGames(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonGamesResponse, error)
	UpdateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error)
	UpdateGames(ctx context.Context, req *pb.CommonGamesRequest) (*pb.CommonIdsResponse, error)
	DeleteGame(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonGameResponse, error)
	DeleteGames(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)

	CreateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error)
	GetMeditation(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeditationResponse, error)
	GetMeditations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonMeditationsResponse, error)
	UpdateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error)
	UpdateMeditations(ctx context.Context, req *pb.CommonMeditationsRequest) (*pb.CommonIdsResponse, error)
	DeleteMeditation(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeditationResponse, error)
	DeleteMeditations(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error)
}
