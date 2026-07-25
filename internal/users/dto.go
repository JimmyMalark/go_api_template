package users

import "time"

type UserResponse struct {
	Xid       string     `json:"xid" example:"9m4e2mr0ui3e8a215n4g"`
	Username  string     `json:"username" example:"john"`
	Email     string     `json:"email" example:"john@example.com"`
	FirstName string     `json:"first_name" example:"john"`
	LastName  string     `json:"last_name" example:"doe"`
	BirthDate *time.Time `json:"birth_date" example:"1998-05-14T00:00:00Z"`
}

type ListUsersResponse struct {
	Items []UserResponse `json:"items"`

	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type CreateUserRequest struct {
	Username  string     `json:"username" validate:"required,min=3,max=100" example:"john"`
	Email     string     `json:"email" validate:"required,email" example:"john@example.com"`
	FirstName string     `json:"first_name" validate:"min=3,max=100" example:"john"`
	LastName  string     `json:"last_name" validate:"min=3,max=100" example:"doe"`
	BirthDate *time.Time `json:"birth_date" validate:"notfuture" example:"1998-05-14T00:00:00Z"`
}
