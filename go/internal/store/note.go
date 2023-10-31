package store

import (
	"context"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id      uint
	Content string
}

func (s *store) GetAllNotes(ctx context.Context) ([]Note, error) {
	var notes []Note
	err := s.db.WithContext(ctx).Find(&notes).Error
	if err != nil {
		return notes, err
	}
	return notes, nil
}

func (s *store) GetNote(ctx context.Context, id uint) (Note, error) {
	var note Note
	err := s.db.WithContext(ctx).First(&note, id).Error
	if err != nil {
		return note, err
	}
	return note, err
}

func (s *store) CreateNote(ctx context.Context, content string) error {
	err := s.db.WithContext(ctx).Create(&Note{
		Content: content,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *store) EditNote(ctx context.Context, id uint, content string) error {
	note := Note{}
	err := s.db.WithContext(ctx).First(&note, id).Update("content", content).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *store) DeleteNote(ctx context.Context, id uint) error {
	note := Note{}
	err := s.db.WithContext(ctx).First(&note, id).Delete(&note).Error
	if err != nil {
		return err
	}
	return nil
}
