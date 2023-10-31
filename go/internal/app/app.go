package app

import (
	"context"
	"main.go/internal/store"
)

type App interface {
	GetAllNotes(ctx context.Context) ([]store.Note, error)
	GetNote(ctx context.Context, id string) (store.Note, error)
	CreateNote(ctx context.Context, content string) error
	EditNote(ctx context.Context, id string, content string) error
	DeleteNote(ctx context.Context, id string) error
}

type app struct {
	store store.Store
}

func New(store store.Store) *app {
	return &app{
		store: store,
	}
}
