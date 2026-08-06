package auth

import "time"

type UserResponse struct {
	Xid       string     `json:"xid" example:"9m4e2mr0ui3e8a215n4g"`
	Username  string     `json:"username" example:"john"`
	Email     string     `json:"email" example:"john@example.com"`
	FirstName string     `json:"first_name" example:"john"`
	LastName  string     `json:"last_name" example:"doe"`
	BirthDate *time.Time `json:"birth_date" example:"1998-05-14T00:00:00Z"`
}

type RegisterRequest struct {
	Username  string     `json:"username" validate:"required,min=3,max=100" example:"john"`
	Email     string     `json:"email" validate:"required,email" example:"john@example.com"`
	FirstName string     `json:"first_name" validate:"min=3,max=100" example:"john"`
	LastName  string     `json:"last_name" validate:"min=3,max=100" example:"doe"`
	Password  string     `validate:"required,min=8"`
	BirthDate *time.Time `json:"birth_date" validate:"notfuture" example:"1998-05-14T00:00:00Z"`
}

type RegisterResult struct {
	User         UserResponse
	SessionToken string
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResult struct {
	User         UserResponse
	SessionToken string
}
