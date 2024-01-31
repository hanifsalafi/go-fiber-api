package repository

import (
	"go-fiber-api/app/database"
	"go-fiber-api/app/database/entity"
	"go-fiber-api/app/module/user/request"
	"go-fiber-api/utils/paginator"
)

type userRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type UserRepository interface {
	GetAll(req request.UserQueryRequest) (Users []*entity.User, paging paginator.Pagination, err error)
	FindOne(id uint) (User *entity.User, err error)
	Create(User *entity.User) (err error)
	Update(id uint, User *entity.User) (err error)
	Delete(id uint) (err error)
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *userRepository) GetAll(req request.UserQueryRequest) (articles []*entity.User, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&entity.User{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&articles).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *userRepository) FindOne(id uint) (User *entity.User, err error) {
	if err := _i.DB.DB.First(&User, id).Error; err != nil {
		return nil, err
	}

	return User, nil
}

func (_i *userRepository) Create(User *entity.User) (err error) {
	return _i.DB.DB.Create(User).Error
}

func (_i *userRepository) Update(id uint, User *entity.User) (err error) {
	return _i.DB.DB.Model(&entity.User{}).
		Where(&entity.User{ID: id}).
		Updates(User).Error
}

func (_i *userRepository) Delete(id uint) error {
	return _i.DB.DB.Delete(&entity.User{}, id).Error
}
