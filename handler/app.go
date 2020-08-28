package handler

import (
	"DisasterManager/dao"
	"DisasterManager/response"
	"net/http"

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

}
