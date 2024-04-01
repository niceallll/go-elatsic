package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"webesapp/logic"
	"webesapp/models"
)

func PostV6CreateQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("create qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.CreateV6Qurey(p); err != nil {
		zap.L().Error("logic  create qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func PostV7CreateQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("create qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.CreateV7Qurey(p); err != nil {
		zap.L().Error("logic  create qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func CreateQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("create qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.CreateQurey(p); err != nil {
		zap.L().Error("logic  create qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func DeleteQurey(c *gin.Context) {
	p := new(models.EsSqlDetail)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.shouldbindjson(p) error", zap.Any("err", err))
		zap.L().Error("delete qurey with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.DeleteQurey(p); err != nil {
		zap.L().Error("logic  delete  qurey failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
