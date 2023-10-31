package store

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store interface {
	GetAllNotes(ctx context.Context) ([]Note, error)
	GetNote(ctx context.Context, id string) (Note, error)
	CreateNote(ctx context.Context, content string) error
	EditNote(ctx context.Context, id string, content string) error
	DeleteNote(ctx context.Context, id string) error
}

type store struct {
	db *gorm.DB
}

func New() *store {
	dsn := "postgres://docker:docker@db:5432/notebook"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Note{})
	return &store{
		db: db,
	}
}
