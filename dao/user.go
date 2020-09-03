package dao

import (
	"DisasterManager/model"
	"DisasterManager/response"
	"DisasterManager/utils"
)

// UserDao 对user模型进行增删查改
type UserDao model.User

// User
func User() *UserDao {
	return &UserDao{}
}

// ReqRegister 注册
type ReqRegister struct {
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Sex         string `json:"sex"`
	Email       string `json:"email" `
	Birthday    string `json:"birthday"`
	Address     string `json:"address"`
	CommunityId int    `json:"communityId" `
	Remark      string `json:"remark"`
}

func (reg *ReqRegister) Register() response.Res {
	u := UserDao{
		UserId:      utils.GetIntID(),
		Username:    reg.Username,
		Phone:       reg.Phone,
		Password:    reg.Password,
		Avatar:      reg.Avatar,
		Sex:         reg.Sex,
		Email:       reg.Email,
		Birthday:    reg.Birthday,
		Address:     reg.Address,
		CommunityId: reg.CommunityId,
		Remark:      reg.Remark,
	}
	if err := model.DB.Table("user").Create(&u).Error; err != nil {
		return response.Res{
			Code: response.CodeDBError,
			Msg:  response.CodeErrMsg[response.CodeDBError],
		}
	}

	return response.Res{
		Code: response.CodeSuccess,
		Data: u,
	}
}

func (u *UserDao) GetUser(userName string) (userList []UserDao, err error) {
	table := model.DB.Table("user")
	if userName != "" {
		table = table.Where("username like ?", "%"+userName+"%")
	}
	if err = table.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}
