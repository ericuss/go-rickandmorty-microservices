package repositories

type RickAndMortyRepository interface {
	RepositoryBase
}

type rickAndMortyRepository struct {
	repositoryBase
}

func NewRickAndMortyRepository() *rickAndMortyRepository {
	return &rickAndMortyRepository{
		repositoryBase: *NewRepositoryBase("Characters"),
	}
}
