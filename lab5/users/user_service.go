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
