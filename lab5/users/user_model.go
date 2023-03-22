package users

type UserModel struct {
	Pk    uint `gorm:"primarykey"`
	Id    string
	Name  string
	Phone string
}

func (UserModel) TableName() string {
	return "user"
}
