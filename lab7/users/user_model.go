package users

type UserModel struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Phone string
}

func (UserModel) TableName() string {
	return "user"
}
