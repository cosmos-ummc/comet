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
		LastDeclared:       patient.LastDeclared,
		SwabCount:          patient.SwabCount,
		Episode:            patient.Episode,
		Type:               patient.Type,
		TypeChangeDate:     patient.TypeChangeDate,
		LastDeclareResult:  patient.LastDeclareResult,
		ExposureDate:       patient.ExposureDate,
		ExposureSource:     patient.ExposureSource,
		DaysSinceExposure:  patient.DaysSinceExposure,
		RegistrationNum:    patient.RegistrationNum,
		AlternateContact:   patient.AlternateContact,
		IsolationAddress:   patient.IsolationAddress,
		SymptomDate:        patient.SymptomDate,
		SwabDate:           patient.SwabDate,
		Remarks:            patient.Remarks,
		Localization:       patient.Localization,
		FeverContDay:       patient.FeverContDay,
		Consent:            patient.Consent,
		PrivacyPolicy:      patient.PrivacyPolicy,
		FeverStartDate:     patient.FeverStartDate,
		DaysSinceSwab:      patient.DaysSinceSwab,
		HomeAddress:        patient.HomeAddress,
		IsSameAddress:      patient.IsSameAddress,
		RegistrationStatus: patient.RegistrationStatus,
	}
}

func PatientToResponse(patient *dto.Patient) *pb.CommonPatientResponse {
	return &pb.CommonPatientResponse{
		Data: &pb.Patient{
			Id:                patient.ID,
			TelegramId:        patient.TelegramID,
			Name:              patient.Name,
			Status:            patient.Status,
			PhoneNumber:       patient.PhoneNumber,
			Email:             patient.Email,
			LastDeclared:      patient.LastDeclared,
			SwabCount:         patient.SwabCount,
			Episode:           patient.Episode,
			Type:              patient.Type,
			TypeChangeDate:    patient.TypeChangeDate,
			LastDeclareResult: patient.LastDeclareResult,
			ExposureDate:      patient.ExposureDate,
			ExposureSource:    patient.ExposureSource,
			DaysSinceExposure: patient.DaysSinceExposure,
			RegistrationNum:   patient.RegistrationNum,
			AlternateContact:  patient.AlternateContact,
			IsolationAddress:  patient.IsolationAddress,
			SymptomDate:       patient.SymptomDate,
			SwabDate:          patient.SwabDate,
			Remarks:           patient.Remarks,
			Localization:      patient.Localization,
			Consent:           patient.Consent,
			PrivacyPolicy:     patient.PrivacyPolicy,
			FeverStartDate:    patient.FeverStartDate,
			FeverContDay:      patient.FeverContDay,
			DaysSinceSwab:     patient.DaysSinceSwab,
			HomeAddress:       patient.HomeAddress,
			IsSameAddress:     patient.IsSameAddress,
		},
	}
}

func PatientsToResponse(patients []*dto.Patient) *pb.CommonPatientsResponse {
	var resps []*pb.Patient
	for _, patient := range patients {
		resp := &pb.Patient{
			Id:                patient.ID,
			TelegramId:        patient.TelegramID,
			Name:              patient.Name,
			Status:            patient.Status,
			PhoneNumber:       patient.PhoneNumber,
			Email:             patient.Email,
			LastDeclared:      patient.LastDeclared,
			SwabCount:         patient.SwabCount,
			Episode:           patient.Episode,
			Type:              patient.Type,
			TypeChangeDate:    patient.TypeChangeDate,
			LastDeclareResult: patient.LastDeclareResult,
			ExposureDate:      patient.ExposureDate,
			ExposureSource:    patient.ExposureSource,
			DaysSinceExposure: patient.DaysSinceExposure,
			RegistrationNum:   patient.RegistrationNum,
			AlternateContact:  patient.AlternateContact,
			IsolationAddress:  patient.IsolationAddress,
			SymptomDate:       patient.SymptomDate,
			SwabDate:          patient.SwabDate,
			Remarks:           patient.Remarks,
			Localization:      patient.Localization,
			CallingStatus:     patient.CallingStatus,
			Consent:           patient.Consent,
			PrivacyPolicy:     patient.PrivacyPolicy,
			FeverStartDate:    patient.FeverStartDate,
			FeverContDay:      patient.FeverContDay,
			DaysSinceSwab:     patient.DaysSinceSwab,
			HomeAddress:       patient.HomeAddress,
			IsSameAddress:     patient.IsSameAddress,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonPatientsResponse{
		Data: resps,
	}

	return rslt
}

// -------------- Patients -----------------

// -------------- Swabs -----------------
func SwabToPb(swab *dto.Swab) *pb.Swab {
	if swab == nil {
		return nil
	}
	return &pb.Swab{
		Id:                  swab.ID,
		PatientId:           swab.PatientID,
		PatientName:         swab.PatientName,
		PatientPhoneNumber:  swab.PatientPhoneNumber,
		Status:              swab.Status,
		Date:                swab.Date,
		Location:            swab.Location,
		IsOtherSwabLocation: swab.IsOtherSwabLocation,
	}
}

func SwabToResponse(swab *dto.Swab) *pb.CommonSwabResponse {
	return &pb.CommonSwabResponse{
		Data: &pb.Swab{
			Id:                  swab.ID,
			PatientId:           swab.PatientID,
			PatientName:         swab.PatientName,
			PatientPhoneNumber:  swab.PatientPhoneNumber,
			Status:              swab.Status,
			Date:                swab.Date,
			Location:            swab.Location,
			IsOtherSwabLocation: swab.IsOtherSwabLocation,
		},
	}
}

func SwabsToResponse(swabs []*dto.Swab) *pb.CommonSwabsResponse {
	var resps []*pb.Swab
	for _, swab := range swabs {
		resp := &pb.Swab{
			Id:                  swab.ID,
			PatientId:           swab.PatientID,
			PatientName:         swab.PatientName,
			PatientPhoneNumber:  swab.PatientPhoneNumber,
			Status:              swab.Status,
			Date:                swab.Date,
			Location:            swab.Location,
			IsOtherSwabLocation: swab.IsOtherSwabLocation,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonSwabsResponse{
		Data: resps,
	}

	return rslt
}

// -------------- Swabs -----------------

// -------------- Activities -----------------
func ActivityToResponse(activity *dto.Activity) *pb.CommonActivityResponse {
	resp := &pb.CommonActivityResponse{
		Data: &pb.Activity{
			Id:         activity.ID,
			UserId:     activity.UserID,
			UserName:   activity.UserName,
			OldPatient: PatientToPb(activity.OldPatient),
			NewPatient: PatientToPb(activity.NewPatient),
			OldSwab:    SwabToPb(activity.OldSwab),
			NewSwab:    SwabToPb(activity.NewSwab),
			Time:       activity.Time,
		},
	}
	return resp
}

func ActivitiesToResponses(activities []*dto.Activity) *pb.CommonActivitiesResponse {
	var resps []*pb.Activity
	for _, activity := range activities {
		resp := &pb.Activity{
			Id:         activity.ID,
			UserId:     activity.UserID,
			UserName:   activity.UserName,
			OldPatient: PatientToPb(activity.OldPatient),
			NewPatient: PatientToPb(activity.NewPatient),
			OldSwab:    SwabToPb(activity.OldSwab),
			NewSwab:    SwabToPb(activity.NewSwab),
			Time:       activity.Time,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonActivitiesResponse{
		Data: resps,
	}

	return rslt
}

// -------------- Activities -----------------

// -------------- Reports -----------------
func CallingReportToResponse(report *dto.CallingReport) *pb.CommonReportResponse {
	resp := &pb.CommonReportResponse{
		Data: &pb.Report{
			DontHaveToCallCount: report.DontHaveToCall,
			PatientCalledCount:  report.PatientCalled,
			UmmcCalledCount:     report.UMMCCalled,
			NoYetCallCount:      report.NoYetCall,
			Date:                report.Date,
		},
	}
	return resp
}

func DeclarationReportToResponse(report *dto.DeclarationReport) *pb.CommonReportResponse {
	resp := &pb.CommonReportResponse{
		Data: &pb.Report{
			Date:            report.Date,
			UndeclaredCount: report.UndeclaredCount,
			DeclaredCount:   report.DeclaredCount,
		},
	}
	return resp
}

func PatientStatusReportToResponse(report *dto.PatientStatusReport) *pb.CommonReportResponse {
	resp := &pb.CommonReportResponse{
		Data: &pb.Report{
			Date:                    report.Date,
			Asymptomatic:            report.Asymptomatic,
			Symptomatic:             report.Symptomatic,
			ConfirmedButNotAdmitted: report.ConfirmedButNotAdmitted,
			ConfirmedAndAdmitted:    report.ConfirmedAndAdmitted,
			Completed:               report.Completed,
			Quit:                    report.Quit,
			Recovered:               report.Recovered,
			PassedAway:              report.PassedAway,
		},
	}
	return resp
}

func CallingReportsToResponse(reports []*dto.CallingReport) *pb.CommonReportsResponse {
	var resp []*pb.Report
	for _, report := range reports {
		r := &pb.Report{
			DontHaveToCallCount: report.DontHaveToCall,
			PatientCalledCount:  report.PatientCalled,
			UmmcCalledCount:     report.UMMCCalled,
			NoYetCallCount:      report.NoYetCall,
			Date:                report.Date,
		}
		resp = append(resp, r)
	}
	return &pb.CommonReportsResponse{Data: resp}
}

func DeclarationReportsToResponse(reports []*dto.DeclarationReport) *pb.CommonReportsResponse {
	var resp []*pb.Report
	for _, report := range reports {
		r := &pb.Report{
			UndeclaredCount: report.UndeclaredCount,
			DeclaredCount:   report.DeclaredCount,
			Date:            report.Date,
		}
		resp = append(resp, r)
	}
	return &pb.CommonReportsResponse{Data: resp}
}

func PatientStatusReportsToResponse(reports []*dto.PatientStatusReport) *pb.CommonReportsResponse {
	var resp []*pb.Report
	for _, report := range reports {
		r := &pb.Report{
			Date:                    report.Date,
			Asymptomatic:            report.Asymptomatic,
			Symptomatic:             report.Symptomatic,
			ConfirmedButNotAdmitted: report.ConfirmedButNotAdmitted,
			ConfirmedAndAdmitted:    report.ConfirmedAndAdmitted,
			Completed:               report.Completed,
			Quit:                    report.Quit,
			Recovered:               report.Recovered,
			PassedAway:              report.PassedAway,
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
			Id:                 user.ID,
			Role:               user.Role,
			DisplayName:        user.DisplayName,
			PhoneNumber:        user.PhoneNumber,
			Email:              user.Email,
			Disabled:           user.Disabled,
			FinalQuestionsJson: user.FinalQuestionsJSON,
			Chart:              user.Chart,
		},
	}
}

func UsersToResponse(users []*dto.User) (*pb.CommonUsersResponse, error) {
	var resps []*pb.User
	for _, user := range users {
		resp := &pb.User{
			Id:          user.ID,
			Role:        user.Role,
			DisplayName: user.DisplayName,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Disabled:    user.Disabled,
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
			Cough:              declaration.Cough,
			Throat:             declaration.Throat,
			Fever:              declaration.Fever,
			Breathe:            declaration.Breathe,
			Chest:              declaration.Chest,
			Blue:               declaration.Blue,
			Drowsy:             declaration.Drowsy,
			HasSymptom:         declaration.HasSymptom,
			SubmittedAt:        declaration.SubmittedAt,
			CallingStatus:      declaration.CallingStatus,
			Date:               declaration.Date,
			PatientName:        declaration.PatientName,
			PatientPhoneNumber: declaration.PatientPhoneNumber,
			DoctorRemarks:      declaration.DoctorRemarks,
			Loss:               declaration.Loss,
		},
	}
}

func DeclarationsToResponse(declarations []*dto.Declaration) *pb.CommonDeclarationsResponse {
	var resps []*pb.Declaration
	for _, declaration := range declarations {
		resp := &pb.Declaration{
			Id:                 declaration.ID,
			PatientId:          declaration.PatientID,
			Cough:              declaration.Cough,
			Throat:             declaration.Throat,
			Fever:              declaration.Fever,
			Breathe:            declaration.Breathe,
			Chest:              declaration.Chest,
			Blue:               declaration.Blue,
			Drowsy:             declaration.Drowsy,
			HasSymptom:         declaration.HasSymptom,
			SubmittedAt:        declaration.SubmittedAt,
			CallingStatus:      declaration.CallingStatus,
			Date:               declaration.Date,
			PatientName:        declaration.PatientName,
			PatientPhoneNumber: declaration.PatientPhoneNumber,
			DoctorRemarks:      declaration.DoctorRemarks,
			Loss:               declaration.Loss,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.CommonDeclarationsResponse{
		Data: resps,
	}
	return rslt
}

// -------------- Declarations -----------------
