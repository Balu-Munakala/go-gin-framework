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

type videoController struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &videoController{
		service: s,
	}
}

func (v *videoController) FindAll() []entity.Video {
	return v.service.FindAll()
}

func (v *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return v.service.Save(video)
	// return video
}
