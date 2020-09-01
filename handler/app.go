package handler

import (
	"DisasterManager/dao"
	"DisasterManager/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register 注册
// @Summary 注册
// @Tags App
// @Param data body dao.ReqRegister true "body"
// @Success 200 {object} response.Res{data=model.User}
// @Router /register [post]
func Register(c *gin.Context) {
	var regDao dao.ReqRegister
	if err := c.ShouldBind(&regDao); err == nil {
		c.JSON(http.StatusOK, regDao.Register())
	} else {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeParamErr,
			Msg:   response.CodeErrMsg[response.CodeParamErr],
			Error: err.Error(),
		})
	}
}

// BillBoard 获取宣传栏
// @Summary 获取宣传栏
// @Tags App
// @Success 200 {object} response.Res{data=model.BillBoard}
// @Router /billboard [get]
func BillBoard(c *gin.Context) {
	d := dao.BillBoard()
	result, err := d.GetBillBoard()
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

// CheckDisaster 检查灾难是否发生
// @Summary 检查灾难是否发生
// @Tags App
// @Param userId query string false "userId"
// @Success 200 {object} response.Res{data=model.Disaster}
// @Router /checkDisaster [get]
func CheckDisaster(c *gin.Context) {
	u, err := strconv.ParseInt(c.Request.FormValue("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	dis := dao.Disaster()
	res, err := dis.GetDisasterByUser(u)
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

// GetRefuge 获取避难所信息
// @Summary 获取避难所信息
// @Tags App
// @Param userId query string false "userId"
// @Success 200 {object} response.Res{data=model.Refuge}
// @Router /getRefuge [get]
func GetRefuge(c *gin.Context) {
	u, err := strconv.ParseInt(c.Request.FormValue("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeDBError,
			Msg:   response.CodeErrMsg[response.CodeDBError],
			Error: err.Error(),
		})
		return
	}
	refuge := dao.Refuge()
	r, err := refuge.GetRefugeByUserId(u)
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
		Data: r,
	})
}

// GetRefugeFacility 获取避难设施信息
// @Summary 获取避难设施信息
// @Tags App
// @Param userId query string false "userId"
// @Success 200 {object} response.Res{data=model.RefugeFacility}
// @Router /getRefugeFacility [get]
func GetRefugeFacility(c *gin.Context) {
	u, err := strconv.ParseInt(c.Request.FormValue("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeParamErr,
			Msg:   response.CodeErrMsg[response.CodeParamErr],
			Error: err.Error(),
		})
		return
	}
	refuge := dao.RefugeFacility()
	r, err := refuge.GetRefugeFacilityByUserId(u)
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
		Data: r,
	})
}

// GetDisasterTypeList 获取灾害类型
// @Summary 获取灾害类型
// @Tags App
// @Success 200 {object} response.Res{data=model.DisasterType}
// @Router /getDisasterTypeList [get]
func GetDisasterTypeList(c *gin.Context) {
	result, err := dao.DisasterType().GetDisasterTypeList()
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

// GetKonwledgeByDisasterType 根据灾害类型获取科普知识
// @Summary 根据灾害类型获取科普知识
// @Tags App
// @Param disasterTypeId query string false "disasterTypeId"
// @Success 200 {object} response.Res{data=model.DisasterType}
// @Router /getkonwledge [get]
func GetKonwledgeByDisasterType(c *gin.Context) {
	disasterType, err := strconv.Atoi(c.Request.FormValue("disasterTypeId"))
	if err != nil {
		c.JSON(http.StatusOK, response.Res{
			Code:  response.CodeParamErr,
			Msg:   response.CodeErrMsg[response.CodeParamErr],
			Error: err.Error(),
		})
		return
	}
	result, err := dao.Knowledge().GetKnowledgeByDisasterType(disasterType)
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
