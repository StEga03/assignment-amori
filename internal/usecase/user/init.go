package user

func New(userRepo userResource) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}
