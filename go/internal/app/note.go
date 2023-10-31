package app

import (
	"context"
	"main.go/internal/store"
)

func (s *app) GetAllNotes(ctx context.Context) ([]store.Note, error) {
	return s.store.GetAllNotes(ctx)
}

func (s *app) GetNote(ctx context.Context, id uint) (store.Note, error) {
	return s.store.GetNote(ctx, id)
}

func (s *app) CreateNote(ctx context.Context, content string) error {
	return s.store.CreateNote(ctx, content)
}

func (s *app) EditNotes(ctx context.Context, id uint, content string) error {
	return s.store.EditNotes(ctx, id, content)
}

func (s *app) DeleteNote(ctx context.Context, id uint) error {
	return s.store.DeleteNote(ctx, id)
}
