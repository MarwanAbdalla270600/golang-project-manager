package user

import "context"

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CreateUser(ctx context.Context, userInput *UserInput) (*UserOutput, error) {
	if err := validateUser(userInput); err != nil {
		return nil, err
	}

	databaseUser, err := s.UserRepository.StoreUser(ctx, userInput)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		Username:  databaseUser.Username,
		Firstname: databaseUser.Firstname,
		Lastname:  databaseUser.Lastname,
	}, nil
}
