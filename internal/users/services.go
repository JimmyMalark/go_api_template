package users

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jimmymalark/go_api_template/internal/cache"
	"github.com/jimmymalark/go_api_template/internal/cache/keys"
	"github.com/jimmymalark/go_api_template/internal/ids"
	"github.com/jimmymalark/go_api_template/internal/pagination"
)

type Service struct {
	repo  UserRepository
	cache cache.Cache
}

func NewService(r UserRepository, c cache.Cache) *Service {
	return &Service{
		repo:  r,
		cache: c,
	}
}

func (s *Service) ListUsers(
	ctx context.Context,
	p pagination.Params,
) (pagination.Response[UserResponse], error) {

	var users pagination.Response[UserResponse]

	if s.cache.Get(ctx, keys.UsersList(p), &users) {
		return users, nil
	}

	users, err := s.loadUsers(ctx, p)
	if err != nil {
		return pagination.Response[UserResponse]{}, err
	}

	_ = s.cache.Set(
		ctx,
		keys.UsersList(p),
		users,
		10*time.Minute,
	)

	return users, nil
}

func (s *Service) getCachedUsers(
	ctx context.Context,
	p pagination.Params,
) (pagination.Response[UserResponse], bool) {
	var res pagination.Response[UserResponse]

	if s.cache.Get(ctx, keys.UsersList(p), &res) {
		return res, true
	}

	return pagination.Response[UserResponse]{}, false
}

func (s *Service) cacheUsers(
	ctx context.Context,
	p pagination.Params,
	res pagination.Response[UserResponse],
) {
	data, err := json.Marshal(res)
	if err != nil {
		return
	}

	_ = s.cache.Set(ctx, keys.UsersList(p), data, 10*time.Minute)
}

func (s *Service) loadUsers(
	ctx context.Context,
	p pagination.Params,
) (pagination.Response[UserResponse], error) {
	if res, ok := s.getCachedUsers(ctx, p); ok {
		return res, nil
	}

	dbUsers, err := s.repo.ListUsers(ctx, p)
	if err != nil {
		return pagination.Response[UserResponse]{}, err
	}

	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return pagination.Response[UserResponse]{}, err
	}

	res := pagination.Response[UserResponse]{
		Items: toUserResponses(dbUsers),
		Pagination: pagination.Metadata{
			Page:       p.Page,
			Limit:      p.Limit,
			TotalItems: int(total),
			TotalPages: (int(total) + p.Limit - 1) / p.Limit,
		},
	}

	s.cacheUsers(ctx, p, res)

	return res, nil
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (UserResponse, error) {
	user, err := s.repo.CreateUser(ctx, toCreateUserParams(req, ids.New()))
	if err != nil {
		return UserResponse{}, err
	}

	// Invalidate the cached list.
	_ = s.cache.DeletePattern(ctx, keys.UsersPrefix+"*")

	return toUserResponse(user), nil
}
