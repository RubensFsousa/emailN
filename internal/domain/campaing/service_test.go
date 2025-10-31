package campaing

import (
	"emailN/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

var (
	request = contract.NewCampaingDTO{
		Name:    "Teste",
		Content: "body",
		Emails:  []string{"email1@e.com"},
	}

	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(request)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	request.Name = ""

	_, err := service.Create(request)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Service_SaveCampaing(t *testing.T) {
	repositoryMock := new(RepositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != request.Name {
			return false
		}
		if campaing.Content != request.Content {
			return false
		}
		if len(campaing.Contacts) != len(request.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(request)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(request)

	assert.Equal("error to save on database", err.Error())
}
