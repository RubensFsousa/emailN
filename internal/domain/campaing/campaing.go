package campaing

import (
	"errors"
	"strings"
	"time"

	"github.com/rs/xid"
)

type Campaing struct {
	ID        string
	CreatedAt time.Time
	Name      string
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}

func NewCampaing(name string, content string, emails []string) (*Campaing, error) {

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}

	if strings.TrimSpace(content) == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("emails is required")
	}

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaing{
		ID:        xid.New().String(),
		CreatedAt: time.Now(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
	}, nil
}
