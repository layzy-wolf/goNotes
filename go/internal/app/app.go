package app

import (
	"context"
	"main.go/internal/store"
)

type App interface {
	GetAllNotes(ctx context.Context) ([]store.Note, error)
	GetNote(ctx context.Context, id uint) (store.Note, error)
	CreateNote(ctx context.Context, content string) error
	EditNotes(ctx context.Context, id uint, content string) error
	DeleteNote(ctx context.Context, id uint) error
}

type app struct {
	store store.Store
}

func New(store store.Store) *app {
	return &app{
		store: store,
	}
}
