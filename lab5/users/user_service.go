package users

type UserService struct {
	repo *Repository
}

func NewUserService(repo *Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) GetAllUsers() ([]User, error) {
	users, err := service.repo.GetAll()
	return users, err
}

func (service *UserService) CreateUser(user User) (string, error) {
	id, err := service.repo.Create(&user)
	return id, err
}

func (service *UserService) RemoveUser(id string) error {
	err := service.repo.Delete(id)
	return err
}
