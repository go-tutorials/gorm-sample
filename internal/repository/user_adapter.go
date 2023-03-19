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

func NewUserAdapter(db *gorm.DB) UserRepository {
	return &UserAdapter{DB: db}
}

type UserAdapter struct {
	DB *gorm.DB
}

func (r *UserAdapter) All(ctx context.Context) (*[]User, error) {
	var users *[]User
	_ = r.DB.Find(&users)
	return users, nil
}

func (r *UserAdapter) Load(ctx context.Context, id string) (*User, error) {
	var user User
	r.DB.First(&user, "id = ?", id)
	return &user, nil
}

func (r *UserAdapter) Create(ctx context.Context, user *User) (int64, error) {
	res := r.DB.Create(&user)
	return res.RowsAffected, nil
}

func (r *UserAdapter) Update(ctx context.Context, user *User) (int64, error) {
	res := r.DB.Save(&user)
	return res.RowsAffected, nil
}

func (r *UserAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	colMap := q.JSONToColumns(user, jsonColumnMap)
	var userModel User
	res := r.DB.Model(&userModel).Where("id = ?", user["id"]).Updates(colMap)
	return res.RowsAffected, nil
}

func (r *UserAdapter) Delete(ctx context.Context, id string) (int64, error) {
	var user User
	res := r.DB.Where("id = ?", id).Delete(&user)
	return res.RowsAffected, nil
}
