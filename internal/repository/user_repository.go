package repository

import (
	"context"
	q "github.com/core-go/sql"
	"gorm.io/gorm"
	"reflect"

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

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
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
	res := r.DB.Save(&user)
	return res.RowsAffected, nil
}

func (r *userRepository) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	colMap := q.JSONToColumns(user, jsonColumnMap)
	var userModel User
	res := r.DB.Model(&userModel).Where("id = ?", user["id"]).Updates(colMap)
	return res.RowsAffected, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) (int64, error) {
	var user User
	res := r.DB.Where("id = ?", id).Delete(&user)
	return res.RowsAffected, nil
}
