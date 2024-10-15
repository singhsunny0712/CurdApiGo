package service

import "errors"

// User struct (Model)
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


// In-memory user "database"
var users = []User{}

func NewUserService() *UserService {
	return &UserService{}
}

// UserService defines the methods for user management
type UserService struct{}

// CreateUser adds a new user to the in-memory "database"
func (s *UserService) CreateUser(name, email string) User {
	newUser := User{
		ID:    len(users) + 1,
		Name:  name,
		Email: email,
	}
	users = append(users, newUser)
	return newUser
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() []User {
	return users
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, errors.New("user not found")
}

// UpdateUser updates the details of a user
func (s *UserService) UpdateUser(id int, name, email string) (User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i].Name = name
			users[i].Email = email
			return users[i], nil
		}
	}
	return User{}, errors.New("user not found")
}

// DeleteUser removes a user by their ID
func (s *UserService) DeleteUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) // Remove user
			return nil
		}
	}
	return errors.New("user not found")
}
