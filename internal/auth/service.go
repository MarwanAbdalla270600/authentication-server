package auth

import (
	"authentication-server/internal/utils"
	"fmt"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Login(data *LoginObject) (*AuthResponse, error)
	Register(data *RegisterObject) (*AuthResponse, error)
}

type service struct {
	repository RepoInterface
}

func NewService(repository RepoInterface) ServiceInterface {
	return &service{repository: repository}
}

func (s *service) Login(data *LoginObject) (*AuthResponse, error) {
	return nil, nil
}

func (s *service) Register(data *RegisterObject) (*AuthResponse, error) {
	fmt.Print(data)
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := UserDAO{
		Id:        uuid.NewString(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashedPassword,
	}

	responseData, err := s.repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: UserDTO{
			Id:        responseData.Id,
			FirstName: responseData.FirstName,
			LastName:  responseData.LastName,
			Email:     responseData.Email,
			CreatedAt: responseData.CreatedAt,
			UpdatedAt: responseData.UpdatedAt,
		},
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}
