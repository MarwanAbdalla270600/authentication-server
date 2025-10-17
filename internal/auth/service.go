package auth

type ServiceInterface interface {
	Login(data LoginObject) (*AuthResponse, error)
	Register(data UserDAO) (*AuthResponse, error)
}

type service struct{}

func NewService() ServiceInterface {
	return &service{}
}

func (s *service) Login(data LoginObject) (*AuthResponse, error) {
	return nil, nil
}

func (s *service) Register(data UserDAO) (*AuthResponse, error) {
	return nil, nil
}



