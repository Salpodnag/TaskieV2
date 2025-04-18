package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID
	Name        string
	Description string
	Color       string
	Privacy     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(name, description string) (*Project, error) {
	if name == "" {
		return nil, fmt.Errorf("project name cannot be empty")
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}
	return &Project{

		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (p *Project) Update(name, description string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	p.Name = name
	p.Description = description
	p.UpdatedAt = time.Now()
	return nil
}
