package user

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CreateUser(userInput *UserInput) (*UserOutput, error) {
	if err := validateUser(userInput); err != nil {
		return nil, err
	}
	// call the repository to store
	// retrieve the user, send it back
	return &UserOutput{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
	}, nil
}
