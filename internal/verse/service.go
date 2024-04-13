package verse

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (me *Service) GetRandomVerse() (*Verse, error) {
	return me.repo.getRandomVerse()
}
