package store

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store interface {
	GetAllNotes(ctx context.Context) ([]Note, error)
	GetNote(ctx context.Context, id uint) (Note, error)
	CreateNote(ctx context.Context, content string) error
	EditNotes(ctx context.Context, id uint, content string) error
	DeleteNote(ctx context.Context, id uint) error
}

type store struct {
	db *gorm.DB
}

func New() *store {
	dsn := "host=localhost user=docker password=docker dbname=docker port=3306 sslmode=disable TimeZone=Asia/Krasnoyarsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Note{})
	return &store{
		db: db,
	}
}
