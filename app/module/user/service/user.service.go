package service

import (
	"go-fiber-api/app/module/user/mapper"
	"go-fiber-api/app/module/user/repository"
	"go-fiber-api/app/module/user/request"
	"go-fiber-api/app/module/user/response"
	"go-fiber-api/utils/paginator"
)

// UserService
type userService struct {
	Repo repository.UserRepository
}

// UserService define interface of IUserService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type UserService interface {
	All(req request.UserQueryRequest) (user []*response.UserResponse, paging paginator.Pagination, err error)
	Show(id uint) (user *response.UserResponse, err error)
	Save(req request.UserCreateRequest) (err error)
	Update(id uint, req request.UserUpdateRequest) (err error)
	Delete(id uint) error
}

// NewArticleService init ArticleService
func NewArticleService(repo repository.UserRepository) UserService {
	return &userService{
		Repo: repo,
	}
}

// All implement interface of ArticleService
func (_i *userService) All(req request.UserQueryRequest) (users []*response.UserResponse, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetAll(req)
	if err != nil {
		return
	}

	for _, result := range results {
		users = append(users, mapper.UserResponseMapper(result))
	}

	return
}

func (_i *userService) Show(id uint) (user *response.UserResponse, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return mapper.UserResponseMapper(result), nil
}

func (_i *userService) Save(req request.UserCreateRequest) (err error) {
	return _i.Repo.Create(req.ToEntity())
}

func (_i *userService) Update(id uint, req request.UserUpdateRequest) (err error) {
	return _i.Repo.Update(id, req.ToEntity())
}

func (_i *userService) Delete(id uint) error {
	return _i.Repo.Delete(id)
}