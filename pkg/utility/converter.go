package utility

import (
	pb "comet/pkg/api"
	"comet/pkg/dto"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// -------------- Patients -----------------
func PatientToPb(patient *dto.Patient) *pb.Patient {
	if patient == nil {
		return nil
	}

	p := &pb.Patient{
		Id:                 patient.ID,
		TelegramId:         patient.TelegramID,
		Name:               patient.Name,
		PhoneNumber:        patient.PhoneNumber,
		Email:              patient.Email,
		IsolationAddress:   patient.IsolationAddress,
		Remarks:            patient.Remarks,
		Consent:            patient.Consent,
		PrivacyPolicy:      patient.PrivacyPolicy,
		HomeAddress:        patient.HomeAddress,
		LastDassTime:       patient.LastDassTime,
		LastIesrTime:       patient.LastIesrTime,
		LastDassResult:     patient.LastDassResult,
		LastIesrResult:     patient.LastIesrResult,
		RegistrationStatus: patient.RegistrationStatus,
		UserId:             patient.UserID,
		DaySinceMonitoring: patient.DaySinceMonitoring,
		HasCompleted:       patient.HasCompleted,
		MentalStatus:       patient.MentalStatus,
		Type:               patient.Type,
		SwabDate:           patient.SwabDate,
		SwabResult:         patient.SwabResult,
		StressStatus:       patient.StressStatus,
		PtsdStatus:         patient.PtsdStatus,
		DepressionStatus:   patient.DepressionStatus,
		AnxietyStatus:      patient.AnxietyStatus,
	}

	// get personality
	resp, err := http.Get(fmt.Sprintf("https://chat.quaranteams.tk/personality?id=%s", patient.ID))
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil && len(string(body)) > 2 {
			bodyString := string(body)
			ids := strings.Split(bodyString[1:len(bodyString)-1], ",")
			i := 0
			for _, rs := range ids {
				ids[i] = rs[1 : len(rs)-1]
				i += 1
			}
			p.Personality = ids
		}
	}

	return p
}

func PatientToResponse(patient *dto.Patient) *pb.CommonPatientResponse {
	p := &pb.Patient{
		Id:                 patient.ID,
		TelegramId:         patient.TelegramID,
		Name:               patient.Name,
		PhoneNumber:        patient.PhoneNumber,
		Email:              patient.Email,
		IsolationAddress:   patient.IsolationAddress,
		Remarks:            patient.Remarks,
		Consent:            patient.Consent,
		PrivacyPolicy:      patient.PrivacyPolicy,
		HomeAddress:        patient.HomeAddress,
		LastDassTime:       patient.LastDassTime,
		LastIesrTime:       patient.LastIesrTime,
		LastDassResult:     patient.LastDassResult,
		LastIesrResult:     patient.LastIesrResult,
		RegistrationStatus: patient.RegistrationStatus,
		UserId:             patient.UserID,
		DaySinceMonitoring: patient.DaySinceMonitoring,
		HasCompleted:       patient.HasCompleted,
		MentalStatus:       patient.MentalStatus,
		Type:               patient.Type,
		SwabDate:           patient.SwabDate,
		SwabResult:         patient.SwabResult,
		StressStatus:       patient.StressStatus,
		PtsdStatus:         patient.PtsdStatus,
		DepressionStatus:   patient.DepressionStatus,
		AnxietyStatus:      patient.AnxietyStatus,
	}

	// get personality
	resp, err := http.Get(fmt.Sprintf("https://chat.quaranteams.tk/personality?id=%s", patient.ID))
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil && len(string(body)) > 2 {
			bodyString := string(body)
			ids := strings.Split(bodyString[1:len(bodyString)-1], ",")
			i := 0
			for _, rs := range ids {
				ids[i] = rs[1 : len(rs)-1]
				i += 1
			}
			p.Personality = ids
		}
	}

	return &pb.CommonPatientResponse{
		Data: p,
	}
}

func PatientsToResponse(patients []*dto.Patient) *pb.CommonPatientsResponse {
	var resps []*pb.Patient
	for _, patient := range patients {

		p := &pb.Patient{
			Id:                 patient.ID,
			TelegramId:         patient.TelegramID,
			Name:               patient.Name,
			PhoneNumber:        patient.PhoneNumber,
			Email:              patient.Email,
			IsolationAddress:   patient.IsolationAddress,
			Remarks:            patient.Remarks,
			Consent:            patient.Consent,
			PrivacyPolicy:      patient.PrivacyPolicy,
			HomeAddress:        patient.HomeAddress,
			LastDassTime:       patient.LastDassTime,
			LastIesrTime:       patient.LastIesrTime,
			LastDassResult:     patient.LastDassResult,
			LastIesrResult:     patient.LastIesrResult,
			RegistrationStatus: patient.RegistrationStatus,
			UserId:             patient.UserID,
			DaySinceMonitoring: patient.DaySinceMonitoring,
			HasCompleted:       patient.HasCompleted,
			MentalStatus:       patient.MentalStatus,
			Type:               patient.Type,
			SwabDate:           patient.SwabDate,
			SwabResult:         patient.SwabResult,
			StressStatus:       patient.StressStatus,
			PtsdStatus:         patient.PtsdStatus,
			DepressionStatus:   patient.DepressionStatus,
			AnxietyStatus:      patient.AnxietyStatus,
		}

		// get personality
		resp, err := http.Get(fmt.Sprintf("https://chat.quaranteams.tk/personality?id=%s", patient.ID))
		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil && len(string(body)) > 2 {
				bodyString := string(body)
				ids := strings.Split(bodyString[1:len(bodyString)-1], ",")
				i := 0
				for _, rs := range ids {
					ids[i] = rs[1 : len(rs)-1]
					i += 1
				}
				p.Personality = ids
			}
		}

		resps = append(resps, p)
	}
	rslt := &pb.CommonPatientsResponse{
		Data: resps,
	}

	return rslt
}

// -------------- Patients -----------------

// -------------- Reports -----------------
func ReportToResponse(report *dto.Report) *pb.CommonReportResponse {
	resp := &pb.CommonReportResponse{
		Data: &pb.Report{
			DepressionNormal:   report.DepressionNormal,
			DepressionMild:     report.DepressionMild,
			DepressionModerate: report.DepressionModerate,
			DepressionSevere:   report.DepressionSevere,
			DepressionExtreme:  report.DepressionExtreme,
			AnxietyNormal:      report.AnxietyNormal,
			AnxietyMild:        report.AnxietyMild,
			AnxietyModerate:    report.AnxietyModerate,
			AnxietySevere:      report.AnxietySevere,
			AnxietyExtreme:     report.AnxietyExtreme,
			StressNormal:       report.StressNormal,
			StressMild:         report.StressMild,
			StressModerate:     report.StressModerate,
			StressSevere:       report.StressSevere,
			StressExtreme:      report.StressExtreme,
			PtsdNormal:         report.PtsdNormal,
			PtsdSevere:         report.PtsdSevere,
			DepressionCounts:   report.DepressionCounts,
			AnxietyCounts:      report.AnxietyCounts,
			StressCounts:       report.StressCounts,
			PtsdCounts:         report.PtsdCounts,
			DepressionStatuses: report.DepressionStatuses,
			AnxietyStatuses:    report.AnxietyStatuses,
			StressStatuses:     report.StressStatuses,
			PtsdStatuses:       report.PtsdStatuses,
		},
	}
	return resp
}

// -------------- Reports -----------------

// -------------- Users -----------------
func UserToPb(user *dto.User) *pb.User {
	u := &pb.User{
		Id:               user.ID,
		Role:             user.Role,
		Name:             user.Name,
		PhoneNumber:      user.PhoneNumber,
		Email:            user.Email,
		BlockList:        user.BlockList,
		Visible:          user.Visible,
		NotFirstTimeChat: user.NotFirstTimeChat,
		InvitedToMeeting: user.InvitedToMeeting,
	}

	return u
}

func UserToResponse(user *dto.User) *pb.CommonUserResponse {
	u := &pb.User{
		Id:               user.ID,
		Role:             user.Role,
		Name:             user.Name,
		PhoneNumber:      user.PhoneNumber,
		Email:            user.Email,
		BlockList:        user.BlockList,
		Visible:          user.Visible,
		NotFirstTimeChat: user.NotFirstTimeChat,
		InvitedToMeeting: user.InvitedToMeeting,
	}

	return &pb.CommonUserResponse{
		Data: u,
	}
}

func UsersToResponse(users []*dto.User) (*pb.CommonUsersResponse, error) {
	var resps []*pb.User
	for _, user := range users {
		u := &pb.User{
			Id:               user.ID,
			Role:             user.Role,
			Name:             user.Name,
			PhoneNumber:      user.PhoneNumber,
			Email:            user.Email,
			BlockList:        user.BlockList,
			Visible:          user.Visible,
			NotFirstTimeChat: user.NotFirstTimeChat,
			InvitedToMeeting: user.InvitedToMeeting,
		}

		resps = append(resps, u)
	}
	rslt := &pb.CommonUsersResponse{
		Data: resps,
	}

	return rslt, nil
}

// -------------- Users -----------------

// -------------- Declarations -----------------
func DeclarationToResponse(declaration *dto.Declaration) *pb.CommonDeclarationResponse {
	return &pb.CommonDeclarationResponse{
		Data: &pb.Declaration{
			Id:               declaration.ID,
			PatientId:        declaration.PatientID,
			Result:           QuestionsToPb(declaration.Result),
			Category:         declaration.Category,
			Score:            declaration.Score,
			SubmittedAt:      declaration.SubmittedAt,
			DoctorRemarks:    declaration.DoctorRemarks,
			Depression:       declaration.Depression,
			Anxiety:          declaration.Anxiety,
			Stress:           declaration.Stress,
			DepressionStatus: declaration.DepressionStatus,
			StressStatus:     declaration.StressStatus,
			AnxietyStatus:    declaration.AnxietyStatus,
			PtsdStatus:       declaration.PtsdStatus,
		},
	}
}

func PbToDeclaration(declaration *pb.Declaration) *dto.Declaration {
	return &dto.Declaration{
		ID:               declaration.Id,
		PatientID:        declaration.PatientId,
		Result:           PbToQuestions(declaration.Result),
		Category:         declaration.Category,
		Score:            declaration.Score,
		SubmittedAt:      declaration.SubmittedAt,
		DoctorRemarks:    declaration.DoctorRemarks,
		Depression:       declaration.Depression,
		Anxiety:          declaration.Anxiety,
		Stress:           declaration.Stress,
		DepressionStatus: declaration.DepressionStatus,
		StressStatus:     declaration.StressStatus,
		AnxietyStatus:    declaration.AnxietyStatus,
		PtsdStatus:       declaration.PtsdStatus,
	}
}

func DeclarationsToResponse(declarations []*dto.Declaration) *pb.CommonDeclarationsResponse {
	var resps []*pb.Declaration
	for _, declaration := range declarations {
		resp := &pb.Declaration{
			Id:               declaration.ID,
			PatientId:        declaration.PatientID,
			Result:           QuestionsToPb(declaration.Result),
			Category:         declaration.Category,
			Score:            declaration.Score,
			SubmittedAt:      declaration.SubmittedAt,
			DoctorRemarks:    declaration.DoctorRemarks,
			Depression:       declaration.Depression,
			Anxiety:          declaration.Anxiety,
			Stress:           declaration.Stress,
			DepressionStatus: declaration.DepressionStatus,
			StressStatus:     declaration.StressStatus,
			AnxietyStatus:    declaration.AnxietyStatus,
			PtsdStatus:       declaration.PtsdStatus,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonDeclarationsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Declarations -----------------

// -------------- Questions -----------------
func QuestionsToPb(questions []*dto.Question) []*pb.Question {
	var resps []*pb.Question
	for _, question := range questions {
		resp := &pb.Question{
			Id:       question.ID,
			Category: question.Category,
			Type:     question.Type,
			Content:  question.Content,
			Score:    question.Score,
		}
		resps = append(resps, resp)
	}
	return resps
}

func PbToQuestions(questions []*pb.Question) []*dto.Question {
	var resps []*dto.Question
	for _, question := range questions {
		resp := &dto.Question{
			ID:       question.Id,
			Category: question.Category,
			Type:     question.Type,
			Content:  question.Content,
			Score:    question.Score,
		}
		resps = append(resps, resp)
	}
	return resps
}

func QuestionToResponse(question *dto.Question) *pb.CommonQuestionResponse {
	return &pb.CommonQuestionResponse{
		Data: &pb.Question{
			Id:       question.ID,
			Category: question.Category,
			Type:     question.Type,
			Content:  question.Content,
			Score:    question.Score,
		},
	}
}

func QuestionsToResponse(questions []*dto.Question) *pb.CommonQuestionsResponse {
	var resps []*pb.Question
	for _, question := range questions {
		resp := &pb.Question{
			Id:       question.ID,
			Category: question.Category,
			Type:     question.Type,
			Content:  question.Content,
			Score:    question.Score,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonQuestionsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Questions -----------------

// -------------- ChatMessages -----------------

func ChatMessageToResponse(chatMessage *dto.ChatMessage) *pb.CommonChatMessageResponse {
	return &pb.CommonChatMessageResponse{
		Data: &pb.ChatMessage{
			Id:        chatMessage.ID,
			RoomId:    chatMessage.RoomID,
			SenderId:  chatMessage.SenderID,
			Content:   chatMessage.Content,
			Timestamp: chatMessage.Timestamp,
		},
	}
}

func ChatMessagesToResponse(chatMessages []*dto.ChatMessage) *pb.CommonChatMessagesResponse {
	var resps []*pb.ChatMessage
	for _, chatMessage := range chatMessages {
		resp := &pb.ChatMessage{
			Id:        chatMessage.ID,
			RoomId:    chatMessage.RoomID,
			SenderId:  chatMessage.SenderID,
			Content:   chatMessage.Content,
			Timestamp: chatMessage.Timestamp,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonChatMessagesResponse{
		Data: resps,
	}
	return rslt
}

// -------------- ChatMessages -----------------

// -------------- ChatRooms -----------------

func ChatRoomToPb(chatRoom *dto.ChatRoom) *pb.ChatRoom {
	return &pb.ChatRoom{
		Id:             chatRoom.ID,
		ParticipantIds: chatRoom.ParticipantIDs,
		Blocked:        chatRoom.Blocked,
		Timestamp:      chatRoom.Timestamp,
		Name:           chatRoom.Name,
	}
}

func ChatRoomToResponse(chatRoom *dto.ChatRoom) *pb.CommonChatRoomResponse {
	return &pb.CommonChatRoomResponse{
		Data: &pb.ChatRoom{
			Id:             chatRoom.ID,
			ParticipantIds: chatRoom.ParticipantIDs,
			Blocked:        chatRoom.Blocked,
			Timestamp:      chatRoom.Timestamp,
			Name:           chatRoom.Name,
		},
	}
}

func ChatRoomsToResponse(chatRooms []*dto.ChatRoom) *pb.CommonChatRoomsResponse {
	var resps []*pb.ChatRoom
	for _, chatRoom := range chatRooms {
		resp := &pb.ChatRoom{
			Id:             chatRoom.ID,
			ParticipantIds: chatRoom.ParticipantIDs,
			Blocked:        chatRoom.Blocked,
			Timestamp:      chatRoom.Timestamp,
			Name:           chatRoom.Name,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonChatRoomsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- ChatRooms -----------------

// -------------- Consultants -----------------

func ConsultantToResponse(consultant *dto.Consultant) *pb.CommonConsultantResponse {
	return &pb.CommonConsultantResponse{
		Data: &pb.Consultant{
			Id:          consultant.ID,
			UserId:      consultant.UserID,
			Name:        consultant.Name,
			PhoneNumber: consultant.PhoneNumber,
			Email:       consultant.Email,
			TakenSlots:  consultant.TakenSlots,
		},
	}
}

func ConsultantsToResponse(consultants []*dto.Consultant) *pb.CommonConsultantsResponse {
	var resps []*pb.Consultant
	for _, consultant := range consultants {
		resp := &pb.Consultant{
			Id:          consultant.ID,
			UserId:      consultant.UserID,
			Name:        consultant.Name,
			PhoneNumber: consultant.PhoneNumber,
			Email:       consultant.Email,
			TakenSlots:  consultant.TakenSlots,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonConsultantsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Consultants -----------------

// -------------- Meetings -----------------

func MeetingToResponse(meeting *dto.Meeting) *pb.CommonMeetingResponse {
	return &pb.CommonMeetingResponse{
		Data: &pb.Meeting{
			Id:           meeting.ID,
			PatientId:    meeting.PatientID,
			ConsultantId: meeting.ConsultantID,
			Time:         meeting.Time,
			Status:       meeting.Status,
		},
	}
}

func MeetingsToResponse(meetings []*dto.Meeting) *pb.CommonMeetingsResponse {
	var resps []*pb.Meeting
	for _, meeting := range meetings {
		resp := &pb.Meeting{
			Id:           meeting.ID,
			PatientId:    meeting.PatientID,
			ConsultantId: meeting.ConsultantID,
			Time:         meeting.Time,
			Status:       meeting.Status,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonMeetingsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Meetings -----------------

// -------------- Feeds -----------------

func FeedToResponse(feed *dto.Feed) *pb.CommonFeedResponse {
	return &pb.CommonFeedResponse{
		Data: &pb.Feed{
			Id:          feed.ID,
			Title:       feed.Title,
			Description: feed.Description,
			Link:        feed.Link,
			ImgPath:     feed.ImgPath,
			Type:        feed.Type,
		},
	}
}

func FeedsToResponse(feeds []*dto.Feed) *pb.CommonFeedsResponse {
	var resps []*pb.Feed
	for _, feed := range feeds {
		resp := &pb.Feed{
			Id:          feed.ID,
			Title:       feed.Title,
			Description: feed.Description,
			Link:        feed.Link,
			ImgPath:     feed.ImgPath,
			Type:        feed.Type,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonFeedsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Feeds -----------------

// -------------- Tips -----------------

func TipToResponse(tip *dto.Tip) *pb.CommonTipResponse {
	return &pb.CommonTipResponse{
		Data: &pb.Tip{
			Id:          tip.ID,
			Title:       tip.Title,
			Description: tip.Description,
		},
	}
}

func TipsToResponse(tips []*dto.Tip) *pb.CommonTipsResponse {
	var resps []*pb.Tip
	for _, tip := range tips {
		resp := &pb.Tip{
			Id:          tip.ID,
			Title:       tip.Title,
			Description: tip.Description,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonTipsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Tips -----------------

// -------------- Games -----------------

func GameToResponse(game *dto.Game) *pb.CommonGameResponse {
	return &pb.CommonGameResponse{
		Data: &pb.Game{
			Id:         game.ID,
			LinkAdr:    game.LinkAdr,
			LinkIos:    game.LinkIos,
			ImgPathAdr: game.ImgPathAdr,
			ImgPathIos: game.ImgPathIos,
		},
	}
}

func GamesToResponse(games []*dto.Game) *pb.CommonGamesResponse {
	var resps []*pb.Game
	for _, game := range games {
		resp := &pb.Game{
			Id:         game.ID,
			LinkAdr:    game.LinkAdr,
			LinkIos:    game.LinkIos,
			ImgPathAdr: game.ImgPathAdr,
			ImgPathIos: game.ImgPathIos,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonGamesResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Games -----------------

// -------------- Meditations -----------------

func MeditationToResponse(meditation *dto.Meditation) *pb.CommonMeditationResponse {
	return &pb.CommonMeditationResponse{
		Data: &pb.Meditation{
			Id:   meditation.ID,
			Link: meditation.Link,
		},
	}
}

func MeditationsToResponse(meditations []*dto.Meditation) *pb.CommonMeditationsResponse {
	var resps []*pb.Meditation
	for _, meditation := range meditations {
		resp := &pb.Meditation{
			Id:   meditation.ID,
			Link: meditation.Link,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonMeditationsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Meditations -----------------
