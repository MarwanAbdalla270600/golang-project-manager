package user

import "errors"

func ValidateUser(user UserInput) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if user.Firstname == "" || user.Lastname == "" {
		return errors.New("firstname and lastname required")
	}

	return nil
}
