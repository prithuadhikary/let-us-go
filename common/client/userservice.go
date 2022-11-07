package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/prithuadhikary/let-us-go/common/model"
)

type UserService interface {
	Retrieve(request *model.RetrieveRequest) (*model.User, error)
}

type userService struct {
	client *resty.Client
}

func (service *userService) Retrieve(request *model.RetrieveRequest) (*model.User, error) {
	response, err := service.client.
		R().
		SetResult(&model.User{}).
		SetBody(request).
		Post("http://localhost:8082/api/users/retrieve")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, nil
	}
	return response.Result().(*model.User), nil
}

func NewUserServiceClient() UserService {
	return &userService{
		client: resty.New(),
	}
}
