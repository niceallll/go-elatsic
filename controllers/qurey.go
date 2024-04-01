package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"webesapp/models"

	//"strconv"
	"webesapp/logic"
)

//func GetV6NginxQurey(c *gin.Context) {
//	// 默认查询，条件查询
//	// 接受查询条件
//	// 判断查询索引内容
//	// 返回查询内容
//	idstr := c.Param("esname")
//	zap.L().Debug(idstr)
//	data, err := logic.GetV6nginxQurey(idstr)
//	if err != nil {
//		zap.L().Error("logic.Query() failed", zap.Error(err))
//	}
//	ResponseSuccess(c, data)
//}

func PostV6Qurey(c *gin.Context) {
	//收集条件查询
	//fmt.Println(data1)

	u := &models.CatNginxStr{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"err": err})
	}
	if len(u.Time) == 0 {
		u.Time = "-20m"
	}
	data, err := logic.PostV6Qurey(u)
	if err != nil {
		zap.L().Error("logic.postnginxqurey() failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}

func PostV7Qurey(c *gin.Context) {
	//收集条件查询
	//fmt.Println(data1)

	u := &models.CatNginxStr{}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"err": err})
	}
	if len(u.Time) == 0 {
		u.Time = "-20m"
	}
	data, err := logic.PostV7Qurey(u)
	if err != nil {
		zap.L().Error("logic.postnginxqurey() failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}

func SqlV6Qurey(c *gin.Context) {
	data, err := logic.SqlV6Qurey()
	if err != nil {
		zap.L().Error("logic sqlv6qurey ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func SqlV7Qurey(c *gin.Context) {
	data, err := logic.SqlV7Qurey()
	if err != nil {
		zap.L().Error("logic sqlv6qurey ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func SSqlV6Qurey(c *gin.Context) {
	data, err := logic.SSqlV6Qurey()
	if err != nil {
		zap.L().Error("logic sqlserver ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//func PostV6CreateQureyHandler(c *gin.Context) {
//	p := new(models.EsSqlDetail)
//	if err := c.ShouldBindJSON(p); err != nil {
//		zap.L().Debug("postv6createqurey error", zap.Any("err", err))
//		zap.L().Debug("create post V6 es with invalid param")
//		return
//	}
//	if err := logic.PostV6CreateQurey(p); err != nil {
//		zap.L().Error("postv6createqurey faild", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//
//}

//func PostV7NginxQurey(c *gin.Context) {
//	//收集条件查询
//	//fmt.Println(data1)
//	var u models.CatNginxStr
//	c.ShouldBind(&u)
//	//if err != nil {
//	//	c.JSON(http.StatusBadRequest, gin.H{
//	//		"error": err.Error(),
//	//	})
//	//} else {
//	//	c.JSON(http.StatusOK, gin.H{
//	//		"status": "ok",
//	//	})
//	//}
//	data, err := logic.PostV7NginxQurey(u.EsName, u.EsType, u.EsMsg, u.EsIndex, u.Time)
//	if err != nil {
//		zap.L().Error("logic.postnginxqurey() failed", zap.Error(err))
//		return
//	}
//	ResponseSuccess(c, data)
//}
