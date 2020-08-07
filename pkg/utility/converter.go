package utility

import (
	pb "comet/pkg/api"
	"comet/pkg/dto"
)

// -------------- Patients -----------------
func PatientToPb(patient *dto.Patient) *pb.Patient {
	if patient == nil {
		return nil
	}
	return &pb.Patient{
		Id:                 patient.ID,
		TelegramId:         patient.TelegramID,
		Name:               patient.Name,
		Status:             patient.Status,
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
	}
}

func PatientToResponse(patient *dto.Patient) *pb.CommonPatientResponse {
	return &pb.CommonPatientResponse{
		Data: &pb.Patient{
			Id:                 patient.ID,
			TelegramId:         patient.TelegramID,
			Name:               patient.Name,
			Status:             patient.Status,
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
		},
	}
}

func PatientsToResponse(patients []*dto.Patient) *pb.CommonPatientsResponse {
	var resps []*pb.Patient
	for _, patient := range patients {
		resp := &pb.Patient{
			Id:                 patient.ID,
			TelegramId:         patient.TelegramID,
			Name:               patient.Name,
			Status:             patient.Status,
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
		}
		resps = append(resps, resp)
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
			Date:   report.Date,
			Counts: report.Counts,
		},
	}
	return resp
}

func ReportsToResponse(reports []*dto.Report) *pb.CommonReportsResponse {
	var resp []*pb.Report
	for _, report := range reports {
		r := &pb.Report{
			Date:   report.Date,
			Counts: report.Counts,
		}
		resp = append(resp, r)
	}
	return &pb.CommonReportsResponse{Data: resp}
}

// -------------- Reports -----------------

// -------------- Users -----------------
func UserToResponse(user *dto.User) *pb.CommonUserResponse {
	return &pb.CommonUserResponse{
		Data: &pb.User{
			Id:          user.ID,
			Role:        user.Role,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			BlockList:   user.BlockList,
		},
	}
}

func UsersToResponse(users []*dto.User) (*pb.CommonUsersResponse, error) {
	var resps []*pb.User
	for _, user := range users {
		resp := &pb.User{
			Id:          user.ID,
			Role:        user.Role,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			BlockList:   user.BlockList,
		}

		resps = append(resps, resp)
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
			Id:                 declaration.ID,
			PatientId:          declaration.PatientID,
			PatientName:        declaration.PatientName,
			PatientPhoneNumber: declaration.PatientPhoneNumber,
			Result:             QuestionsToPb(declaration.Result),
			Category:           declaration.Category,
			Score:              declaration.Score,
			Status:             declaration.Status,
			SubmittedAt:        declaration.SubmittedAt,
			DoctorRemarks:      declaration.DoctorRemarks,
		},
	}
}

func PbToDeclaration(declaration *pb.Declaration) *dto.Declaration {
	return &dto.Declaration{
		ID:                 declaration.Id,
		PatientID:          declaration.PatientId,
		PatientName:        declaration.PatientName,
		PatientPhoneNumber: declaration.PatientPhoneNumber,
		Result:             PbToQuestions(declaration.Result),
		Category:           declaration.Category,
		Score:              declaration.Score,
		Status:             declaration.Status,
		SubmittedAt:        declaration.SubmittedAt,
		DoctorRemarks:      declaration.DoctorRemarks,
	}
}

func DeclarationsToResponse(declarations []*dto.Declaration) *pb.CommonDeclarationsResponse {
	var resps []*pb.Declaration
	for _, declaration := range declarations {
		resp := &pb.Declaration{
			Id:                 declaration.ID,
			PatientId:          declaration.PatientID,
			PatientName:        declaration.PatientName,
			PatientPhoneNumber: declaration.PatientPhoneNumber,
			Result:             QuestionsToPb(declaration.Result),
			Category:           declaration.Category,
			Score:              declaration.Score,
			Status:             declaration.Status,
			SubmittedAt:        declaration.SubmittedAt,
			DoctorRemarks:      declaration.DoctorRemarks,
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
