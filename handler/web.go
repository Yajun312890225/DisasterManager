package handler

import (
	"DisasterManager/dao"
	"DisasterManager/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCommunities 获取社区
// @Summary 获取社区
// @Tags Web
// @Success 200 {object} response.Res{data=model.Community}
// @Router /getCommunities [get]
func GetCommunities(c *gin.Context) {
	d := dao.Community()
	result, err := d.GetAllCommunity()
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Res{
		Code: response.CodeSuccess,
		Data: result,
	})
}

// GetDisaster 获取社区受灾
// @Summary 获取社区受灾
// @Tags Web
// @Param communityId query string false "communityId"
// @Success 200 {object} response.Res{data=model.Disaster}
// @Router /getDisaster [get]
func GetDisaster(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.FormValue("communityId"))
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeParamErr,
			Msg:   response.CodeErrMsg[response.CodeParamErr],
			Error: err.Error(),
		})
		return
	}
	dis := dao.Disaster()
	res, err := dis.GetDisasterByCommunityId(id)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Res{
		Code: response.CodeSuccess,
		Data: res,
	})
}

// GetHistoryDisasterDetail 获取社区灾害历史位置信息
// @Summary 获取社区灾害历史位置信息
// @Tags Web
// @Param communityId query string false "communityId"
// @Param userId query string false "userId"
// @Success 200 {object} response.Res{data=model.Disaster}
// @Router /getHistoryDisasterDetail [get]
func GetHistoryDisasterDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.FormValue("communityId"))
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	var userId int64
	if u := c.Request.FormValue("userId"); u != "" {
		userId, err = strconv.ParseInt(u, 10, 64)
	}
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeParamErr,
			Msg:   response.CodeErrMsg[response.CodeParamErr],
			Error: err.Error(),
		})
		return
	}
	location := dao.Location()
	res, err := location.GetLocationByDisasterId(id, userId)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Res{
		Code: response.CodeSuccess,
		Data: res,
	})
}

// GetUser 获取用户
// @Summary 获取用户
// @Tags Web
// @Param userName query string false "userName"
// @Success 200 {object} response.Res{data=model.Disaster}
// @Router /getUser [get]
func GetUser(c *gin.Context) {
	userName := c.Request.FormValue("userName")

	u := dao.User()
	res, err := u.GetUser(userName)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Res{
		Code: response.CodeSuccess,
		Data: res,
	})
}
