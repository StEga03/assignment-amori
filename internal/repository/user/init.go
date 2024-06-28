package user

func New(db databaseResource, sf sonyFlakeResource, jwtSetupKey string) *Repository {
	return &Repository{
		db:          db,
		sf:          sf,
		jwtSetupKey: jwtSetupKey,
	}
}
