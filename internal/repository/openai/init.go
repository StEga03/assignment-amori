package openai

func New(openaiService OpenaiResources) *Repository {
	return &Repository{
		openaiService: openaiService,
	}
}
