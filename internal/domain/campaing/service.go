package campaing

import (
	"emailN/internal/contract"
	"errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(request contract.NewCampaingDTO) (string, error) {

	camping, err := NewCampaing(request.Name, request.Content, request.Emails)
	if err != nil {
		return "", errors.New(err.Error())
	}
	err = s.Repository.Save(camping)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return camping.ID, nil
}
