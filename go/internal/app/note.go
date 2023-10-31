package app

import (
	"context"
	"main.go/internal/store"
)

func (s *app) GetAllNotes(ctx context.Context) ([]store.Note, error) {
	return s.store.GetAllNotes(ctx)
}

func (s *app) GetNote(ctx context.Context, id string) (store.Note, error) {
	return s.store.GetNote(ctx, id)
}

func (s *app) CreateNote(ctx context.Context, content string) error {
	return s.store.CreateNote(ctx, content)
}

func (s *app) EditNote(ctx context.Context, id string, content string) error {
	return s.store.EditNote(ctx, id, content)
}

func (s *app) DeleteNote(ctx context.Context, id string) error {
	return s.store.DeleteNote(ctx, id)
}
