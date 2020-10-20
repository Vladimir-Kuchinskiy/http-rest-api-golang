package users

// ServiceI ...
type ServiceI interface {
	Create(*CreateDTO) (*User, error)
}

// Service ...
type Service struct {
	usersRepository RepositoryI
}

// NewService ...
func NewService(usersRepository RepositoryI) *Service {
	return &Service{
		usersRepository: usersRepository,
	}
}

// CreateDTO ...
type CreateDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create ...
func (s *Service) Create(params *CreateDTO) (*User, error) {
	u := &User{
		Email:    params.Email,
		Password: params.Password,
	}
	if err := s.usersRepository.Create(u); err != nil {
		return nil, err
	}

	u.Sanitize()

	return u, nil
}
