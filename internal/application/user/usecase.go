package user

type Usercase struct {
	repo UserRepository
}

func NewUsercase(repo UserRepository) *Usercase {
	return &Usercase{
		repo: repo,
	}
}
