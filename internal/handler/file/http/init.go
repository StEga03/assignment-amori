package http

func New(fileUC FileUC) *Handler {
	return &Handler{
		fileUC: fileUC,
	}
}
