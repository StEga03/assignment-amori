package http

func New(UserUC UserUsecase) *Handler {
	return &Handler{
		UserUC: UserUC,
	}
}
