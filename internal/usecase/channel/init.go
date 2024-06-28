package channel

func New(
	consistency consistencyResource,
	channelRepo channelResource,
	messageRepo messageResource,
	openaiRepo openaiResource,
	userRepo userResource,
	sf sonyFlakeResource,
) *Usecase {
	return &Usecase{
		consistency: consistency,
		channelRepo: channelRepo,
		messageRepo: messageRepo,
		openaiRepo:  openaiRepo,
		userRepo:    userRepo,
		sf:          sf,
	}
}
