package http

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/app"
)

type Api struct {
	engine *gin.Engine
	app    app.App
}

func New(app app.App) *Api {
	a := &Api{
		engine: gin.New(),
		app:    app,
	}
	a.SetupRoutes()
	return a
}

func (s Api) SetupRoutes() {
	s.engine.Use(gin.Recovery())

	notes := s.engine.Group("/notes")
	{
		notes.GET("", s.GetAllNotes)
		notes.GET("/:id", s.GetNote)
		notes.GET("/delete/:id", s.DeleteNote)
		notes.POST("", s.CreateNote)
		notes.POST("/edit/:id", s.EditNote)
	}
}

func (s *Api) Engine() {
	s.engine.Run()
}
