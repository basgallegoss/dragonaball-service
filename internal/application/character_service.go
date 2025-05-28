package application

import "github.com/basgallegoss/dragonball-service/internal/domain"

type CharacterService struct {
	repo     domain.Repository
	external ExternalAPI
}

type ExternalAPI interface {
	FetchByName(name string) (domain.Character, error)
}

func NewCharacterService(r domain.Repository, e ExternalAPI) *CharacterService {
	return &CharacterService{repo: r, external: e}
}

func (s *CharacterService) GetOrCreate(name string) (domain.Character, error) {
	if c, err := s.repo.FindByName(name); err == nil && c != nil {
		return *c, nil
	}

	fetched, err := s.external.FetchByName(name)
	if err != nil {
		return domain.Character{}, err
	}
	if err := s.repo.Save(fetched); err != nil {
		return domain.Character{}, err
	}
	return fetched, nil
}
