package auth

import (
	"authentication-server/internal/entity"
	"authentication-server/internal/utils"
	"errors"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Login(data *entity.LoginObject) (*entity.AuthResponse, error)
	Register(data *entity.RegisterObject) (*entity.AuthResponse, error)
}

type service struct {
	repository RepoInterface
}

func NewService(repository RepoInterface) ServiceInterface {
	return &service{repository: repository}
}

func (s *service) Login(data *entity.LoginObject) (*entity.AuthResponse, error) {
	foundUser, err := s.repository.GetUserByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	if !utils.ComparePassword(foundUser.Password, data.Password) {
		return nil, errors.New("password is incorrect")
	}

	accessToken, err := utils.CreateAccessToken(foundUser)
	//refreshToken := uuid.NewString()
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		User: entity.UserDTO{
			Id:        foundUser.Id,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Email:     foundUser.Email,
			Role:      foundUser.Role,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
		},
		AccessToken: accessToken,
		//RefreshToken: refreshToken,
	}, nil
}

func (s *service) Register(data *entity.RegisterObject) (*entity.AuthResponse, error) {
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := entity.UserDAO{
		Id:        uuid.NewString(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashedPassword,
		Role:      "user",
	}

	responseData, err := s.repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	accessToken, err := utils.CreateAccessToken(responseData)
	//refreshToken := uuid.NewString()
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		User: entity.UserDTO{
			Id:        responseData.Id,
			FirstName: responseData.FirstName,
			LastName:  responseData.LastName,
			Email:     responseData.Email,
			Role:      responseData.Role,
			CreatedAt: responseData.CreatedAt,
			UpdatedAt: responseData.UpdatedAt,
		},
		AccessToken: accessToken,
		//RefreshToken: refreshToken,
	}, nil
}
