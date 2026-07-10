package services

import (
	"errors"
	"fmt"
	"strings"

	"todolist-api/models"
	"todolist-api/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) GetByID(id int64) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("id must be a positive integer")
	}

	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	if todo == nil {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (s *TodoService) Create(title, description string) (*models.Todo, error) {
	if err := validateTitle(title); err != nil {
		return nil, err
	}
	if err := validateDescription(description); err != nil {
		return nil, err
	}

	todo := &models.Todo{
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Completed:   false,
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	// Reload from DB to get created_at, updated_at
	created, err := s.repo.FindByID(todo.ID)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return created, nil
}

func (s *TodoService) Update(id int64, title, description string, completed bool) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("id must be a positive integer")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	if existing == nil {
		return nil, errors.New("todo not found")
	}

	if err := validateTitle(title); err != nil {
		return nil, err
	}
	if err := validateDescription(description); err != nil {
		return nil, err
	}

	todo := &models.Todo{
		ID:          id,
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Completed:   completed,
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return updated, nil
}

func (s *TodoService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("id must be a positive integer")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}
	if existing == nil {
		return errors.New("todo not found")
	}

	return s.repo.Delete(id)
}

func (s *TodoService) MarkComplete(id int64) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("id must be a positive integer")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	if existing == nil {
		return nil, errors.New("todo not found")
	}

	if err := s.repo.MarkComplete(id); err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return updated, nil
}

func (s *TodoService) MarkIncomplete(id int64) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("id must be a positive integer")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	if existing == nil {
		return nil, errors.New("todo not found")
	}

	if err := s.repo.MarkIncomplete(id); err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return updated, nil
}

func validateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("title is required")
	}
	if len(title) < 3 {
		return errors.New("title must be at least 3 characters")
	}
	if len(title) > 255 {
		return errors.New("title must not exceed 255 characters")
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 1000 {
		return errors.New("description must not exceed 1000 characters")
	}
	return nil
}