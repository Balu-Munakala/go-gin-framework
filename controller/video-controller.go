package controller

import (
	"github.com/Balu-Munakala/go-gin-framework/entity"
	"github.com/Balu-Munakala/go-gin-framework/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &controller{service: s}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return c.service.Save(video)
	// return video
}
