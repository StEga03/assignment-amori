package http

func New(ChannelUC ChannelUsecase) *Handler {
	return &Handler{
		ChannelUC: ChannelUC,
	}
}
