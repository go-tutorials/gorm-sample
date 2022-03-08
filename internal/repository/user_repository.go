package repository

import (
	"context"
	"gorm.io/gorm"

	. "go-service/internal/model"
)

type UserRepository interface {
	All(ctx context.Context) (*[]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewUserRepository(gormDb *gorm.DB) UserRepository {
	return &userRepository{DB: gormDb}
}

type userRepository struct {
	DB *gorm.DB
}

func (r *userRepository) All(ctx context.Context) (*[]User, error) {
	var users *[]User
	_ = r.DB.Find(&users)
	return users, nil
}

func (r *userRepository) Load(ctx context.Context, id string) (*User, error) {
	var user User
	r.DB.First(&user, "id = ?", id)
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *User) (int64, error) {
	res := r.DB.Create(&user)
	return res.RowsAffected, nil
}

func (r *userRepository) Update(ctx context.Context, user *User) (int64, error) {
	res := r.DB.Model(&user).Updates(User{Username: user.Username, Email: user.Email, Phone: user.Phone, DateOfBirth: user.DateOfBirth})
	return res.RowsAffected, nil
}

func (r *userRepository) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	var userModel User
	res := r.DB.Model(&userModel).Where("id = ?", user["id"]).Updates(user)
	return res.RowsAffected, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) (int64, error) {
	var user User
	res := r.DB.Where("id = ?", id).Delete(&user)
	return res.RowsAffected, nil
}
