package users

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetById(id string) (*User, error) {
	userModel := &UserModel{}

	err := r.db.First(userModel, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	user := ToUser(userModel)
	return &user, nil
}

func (r *Repository) GetAll() ([]User, error) {
	var models []UserModel
	err := r.db.Find(&models).Error

	if err != nil {
		return nil, err
	}

	users := make([]User, len(models))

	for i := 0; i < len(models); i++ {
		users[len(models)-i-1] = ToUser(&models[i])
	}

	return users, nil
}

func (r *Repository) Create(user *User) (string, error) {
	if len(user.Id) == 0 {
		user.Id = uuid.New().String()
	}
	model := ToModel(user)
	return user.Id, r.db.Create(&model).Error
}

func (r *Repository) Update(user *User) error {
	model := ToModel(user)
	return r.db.Model(&UserModel{}).Where("id = ?", model.Id).Updates(model).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&UserModel{}, "id = ?", id).Error
}

func (r *Repository) Clear() error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE %s;", UserModel{}.TableName())).Error
}

func ToUser(model *UserModel) User {
	user := User{
		Id:    model.Id,
		Name:  model.Name,
		Phone: model.Phone,
	}

	return user
}

func ToModel(user *User) UserModel {
	model := UserModel{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}

	return model
}
