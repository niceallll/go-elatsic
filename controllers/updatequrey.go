package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"webesapp/logic"
	"webesapp/models"
)

func PostV6UpdateQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("create qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.UpdateV6Qurey(p); err != nil {
		zap.L().Error("logic  create qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func UpdateQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("create qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.UpdateQurey(p); err != nil {
		zap.L().Error("logic  create qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
