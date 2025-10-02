package campaing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "X"
	content = "body"
	emails  = []string{"email1@a.com", "email2@a.com"}
)

func Test_NewCampaing_CreateNewCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, emails)

	assert.NotNil(campaing.ID)
	assert.NotNil(campaing.CreatedAt)
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(emails))
}

func Test_NewCampaing_NameIsRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, emails)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaing_ContentIsRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", emails)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaing_EmailsIsRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{})

	assert.Equal("emails is required", err.Error())
}
