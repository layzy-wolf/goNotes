package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type note struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
}

func (s *Api) GetAllNotes(c *gin.Context) {
	notes, err := s.app.GetAllNotes(c)
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	var noteList []gin.H
	for _, note := range notes {
		noteList = append(noteList, gin.H{
			"id":      note.Id,
			"content": note.Content,
		})
	}
	c.JSON(http.StatusOK, noteList)
}

func (s *Api) GetNote(c *gin.Context) {
	var note note
	if err := c.ShouldBindUri(&note.Id); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	noteList, err := s.app.GetNote(c, note.Id)
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      noteList.Id,
		"content": noteList.Content,
	})
}

func (s *Api) CreateNote(c *gin.Context) {
	newNote := &note{}
	if err := c.BindJSON(newNote); err != nil {
		logrus.Println(err.Error())
		return
	}

	err := s.app.CreateNote(c, newNote.Content)
	if err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}

func (s *Api) EditNotes(c *gin.Context) {
	var noteItem note
	if err := c.ShouldBindUri(&noteItem.Id); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	note := &note{}
	if err := c.BindJSON(note); err != nil {
		logrus.Println(err.Error())
		return
	}

	if err := s.app.EditNotes(c, noteItem.Id, note.Content); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully",
	})
}

func (s *Api) DeleteNote(c *gin.Context) {
	var noteItem note
	if err := c.ShouldBindUri(&noteItem.Id); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if err := s.app.DeleteNote(c, noteItem.Id); err != nil {
		logrus.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully",
	})
}
