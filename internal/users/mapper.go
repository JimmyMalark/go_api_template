package users

import (
	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
)

func toUserResponse(user db.User) UserResponse {
	return UserResponse{
		Xid:       user.Xid,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
	}
}

func toUserResponses(dbUsers []db.User) []UserResponse {
	users := make([]UserResponse, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = toUserResponse(u)
	}

	return users
}

func toCreateUserParams(req CreateUserRequest, xid string) db.CreateUserParams {
	return db.CreateUserParams{
		Xid:       xid,
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: req.BirthDate,
	}
}
