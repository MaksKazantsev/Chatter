package validator

import (
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

type Validator interface {
	ValidateRegReq(req *pkg.RegisterReq) error
	ValidateLogReq(req *pkg.LoginReq) error
	ValidateSendCodeReq(req *pkg.SendReq) error
	ValidateVerifyCodeReq(req *pkg.VerifyReq) error
	ValidateRecoveryReq(req *pkg.RecoveryReq) error
}

func NewValidator() Validator {
	return &validator{
		regExpEmail: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	}
}

type validator struct {
	regExpEmail *regexp.Regexp
}

func (v validator) ValidateRecoveryReq(req *pkg.RecoveryReq) error {
	ok := v.regExpEmail.MatchString(req.Email)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid email")
	}
	if len(req.Password) < 5 || len(req.Password) > 25 {
		return status.Error(codes.InvalidArgument, "Password can't be shorter than 5 or longer than 25 symbols")
	}
	return nil
}

func (v validator) ValidateVerifyCodeReq(req *pkg.VerifyReq) error {
	ok := v.regExpEmail.MatchString(req.Email)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid email")
	}
	if len(req.Code) != 4 {
		return status.Error(codes.InvalidArgument, "Code must be 4-symbol")
	}
	return nil
}

func (v validator) ValidateSendCodeReq(req *pkg.SendReq) error {
	ok := v.regExpEmail.MatchString(req.Email)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid email")
	}
	return nil
}

func (v validator) ValidateLogReq(req *pkg.LoginReq) error {
	ok := v.regExpEmail.MatchString(req.Email)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid email")
	}
	if len(req.Password) < 5 || len(req.Password) > 25 {
		return status.Error(codes.InvalidArgument, "Password can't be shorter than 5 or longer than 25 symbols")
	}
	return nil
}

func (v validator) ValidateRegReq(req *pkg.RegisterReq) error {
	ok := v.regExpEmail.MatchString(req.Email)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid email")
	}
	if len(req.Password) < 5 || len(req.Password) > 25 {
		return status.Error(codes.InvalidArgument, "Password can't be shorter than 5 or longer than 25 symbols")
	}
	if len(req.Username) < 2 || len(req.Username) > 15 {
		return status.Error(codes.InvalidArgument, "Username can't be shorter than 2 or longer than 15 symbols")
	}
	return nil
}
