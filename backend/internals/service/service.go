package service

import "github.com/marcofilho/go-url-shortner/backend/internals/database"

type Service struct {
	Repo database.Repository
}

func NewService(repo database.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Save(url, id string) error {
	return s.Repo.SaveUrl(id, url)
}

func (s *Service) Get(id string) (string, error) {
	return s.Repo.GetUrl(id)
}
