package constants

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidArgumentError       = status.Error(codes.InvalidArgument, "Invalid input.")
	InvalidPhoneNumberError    = status.Error(codes.InvalidArgument, "Invalid phone number.")
	InvalidDateError           = status.Error(codes.InvalidArgument, "Invalid date.")
	InvalidPatientStatusError  = status.Error(codes.InvalidArgument, "Invalid patient status.")
	InvalidAddressError        = status.Error(codes.InvalidArgument, "Invalid address.")
	InvalidLanguageError       = status.Error(codes.InvalidArgument, "Invalid language.")
	InvalidQuestionStatusError = status.Error(codes.InvalidArgument, "Invalid swab status.")
	InvalidPatientNameError    = status.Error(codes.InvalidArgument, "Invalid patient name.")
	InvalidPatientIDError      = status.Error(codes.InvalidArgument, "Invalid patientID.")
	InvalidPatientTypeError    = status.Error(codes.InvalidArgument, "Invalid patient type.")
	InvalidEmailError          = status.Error(codes.InvalidArgument, "Invalid email.")
	InvalidRoleError           = status.Error(codes.InvalidArgument, "Invalid role.")
	InvalidPasswordError       = status.Error(codes.InvalidArgument, "Invalid password, please ensure that password is more than 6 characters.")
	RemarksTooLongError        = status.Error(codes.InvalidArgument, "Remarks too long.")

	PatientAlreadyExistError     = status.Error(codes.AlreadyExists, "Patient already exist!")
	QuestionAlreadyExistError    = status.Error(codes.AlreadyExists, "Question already exist!")
	UserAlreadyExistError        = status.Error(codes.AlreadyExists, "User already exist!")
	DeclarationAlreadyExistError = status.Error(codes.AlreadyExists, "Declaration already exist!")
	PhoneNumberAlreadyExistError = status.Error(codes.AlreadyExists, "Phone number already exist, please use another phone number.")
	EmailAlreadyExistError       = status.Error(codes.AlreadyExists, "Email already exist, please use another email.")
	ChatMessageAlreadyExistError = status.Error(codes.AlreadyExists, "Chat message already exist!")
	ChatRoomAlreadyExistError    = status.Error(codes.AlreadyExists, "Chat room already exist!")
	ConsultantAlreadyExistError  = status.Error(codes.AlreadyExists, "Consultant already exist!")
	MeetingAlreadyExistError     = status.Error(codes.AlreadyExists, "Meeting already exist!")

	UserNotFoundError        = status.Error(codes.NotFound, "User not found!")
	PatientNotFoundError     = status.Error(codes.NotFound, "Patient not found!")
	ActivityNotFoundError    = status.Error(codes.NotFound, "Activity not found!")
	QuestionNotFoundError    = status.Error(codes.NotFound, "Question not found!")
	DeclarationNotFoundError = status.Error(codes.NotFound, "Declaration not found!")
	MetadataNotFoundError    = status.Error(codes.NotFound, "Metadata not found!")
	ChatMessageNotFoundError = status.Error(codes.NotFound, "ChatMessage not found!")
	ChatRoomNotFoundError    = status.Error(codes.NotFound, "ChatRoom not found!")
	ConsultantNotFoundError  = status.Error(codes.NotFound, "Consultant not found!")
	MeetingNotFoundError     = status.Error(codes.NotFound, "Meeting not found!")

	UserOperationError         = status.Error(codes.Internal, "Authentication Service failed. Might be due to invalid input.")
	UnauthorizedAccessError    = status.Error(codes.Unauthenticated, "User is not authorized to perform this action!")
	PasswordResetLimitError    = status.Error(codes.Internal, "Password reset limit exceeded!")
	InvalidPasswordVerifyError = status.Error(codes.InvalidArgument, "Invalid password.")
	CreateTokenFailedError     = status.Error(codes.Internal, "Failed to create token.")
	VerifyTokenFailedError     = status.Error(codes.Internal, "Failed to verify token.")

	OperationUnsupportedError = status.Error(codes.Unimplemented, "Operation unsupported.")
	InternalError             = status.Error(codes.Internal, "Server unavailable, please try again.")
)
