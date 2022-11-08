package service

import (
	"github.com/prithuadhikary/let-us-go/common/auth"
	"github.com/prithuadhikary/let-us-go/common/client"
	"github.com/prithuadhikary/let-us-go/common/errors"
	"github.com/prithuadhikary/let-us-go/common/model"
	"net/http"
)

type LoginService interface {
	Login(request *model.LoginRequest) (*model.LoginResponse, errors.ServiceError)
}

type loginService struct {
	userService  client.UserService
	tokenService auth.TokenService
}

func (service *loginService) Login(request *model.LoginRequest) (*model.LoginResponse, errors.ServiceError) {
	// Retrieve the user for the
	user, clientErr := service.userService.Retrieve(&model.RetrieveRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if clientErr != nil {
		return nil, errors.InternalServerError()
	}
	// If user not found, respond with unauthorised and error message.
	if user == nil {
		return nil, errors.NewServiceError(
			"login.failure",
			"Username or password or both incorrect.",
			http.StatusUnauthorized,
		)
	}
	// Else generate a JWT token with standard and private claims(only role for now).
	token, err := service.tokenService.Generate(user)
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{
		Token: token,
	}, nil
}

func NewLoginService(userService client.UserService, tokenService auth.TokenService) LoginService {
	return &loginService{
		userService:  userService,
		tokenService: tokenService,
	}
}
