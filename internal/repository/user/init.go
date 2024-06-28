package user

func New(db databaseResource, sf sonyFlakeResource) *Repository {
	return &Repository{
		db: db,
		sf: sf,
	}
}
