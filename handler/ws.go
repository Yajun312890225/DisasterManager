package handler

import (
	"DisasterManager/dao"
	"DisasterManager/model"
	"DisasterManager/pkg/ws"
	"DisasterManager/utils"
	"encoding/json"
	"log"
	"time"

	"gopkg.in/olahol/melody.v1"
)

type MessageData struct {
	CMDID      int         `json:"cmdId"`
	DeviceType int         `json:"deviceType"` //0 app 1 web
	Data       interface{} `json:"data"`
}
type NotifyMessageData struct {
	CMDID int         `json:"cmdId"`
	Data  interface{} `json:"data"`
}

type RspMessageData struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

func Connect(s *melody.Session) {
	// fmt.Println("Connect")
	ID := utils.GetID()
	s.Set("cid", ID)
	ws.GetSessionMaster().SetSession(ID, s)

}

func DisConnect(s *melody.Session) {
	// fmt.Println("DisConnect")
	cid, exist := s.Get("cid")
	if exist {
		ws.GetSessionMaster().DelSessionByCID(cid.(string))
	}
}

func Message(s *melody.Session, msg []byte) {
	var data MessageData
	json.Unmarshal(msg, &data)
	go ConnHandle(data, s)
}

func ConnHandle(data MessageData, s *melody.Session) {
	switch data.CMDID {
	case model.OnLogin:
		s.Set("type", data.DeviceType)
		byteData, err := json.Marshal(RspMessageData{Code: 0, Msg: "登录成功"})
		if err != nil {
			log.Println(err)
		}
		go s.Write(byteData)
	case model.OnLocate:

		d := data.Data.(map[string]interface{})
		location := dao.Location()
		userId, exist := d["userId"]
		if exist {
			location.UserId = int64(userId.(float64))
		}
		disa, exist := d["disasterId"]
		if exist {
			location.DisasterId = int(disa.(float64))
		}
		lo, exist := d["longitude"]
		if exist {
			location.Longitude = lo.(float64)
		}
		la, exist := d["latitude"]
		if exist {
			location.Latitude = la.(float64)
		}
		if err := location.InsertLocation(); err != nil {
			log.Println(err)
			return
		}
		d["createdAt"] = time.Now()
		byteData, err := json.Marshal(NotifyMessageData{CMDID: model.NotifyLocate, Data: d})
		if err != nil {
			log.Println(err)
		}
		go ws.GetSessionMaster().WriteToWeb(byteData)

		byteData, err = json.Marshal(RspMessageData{Code: 0, Msg: "上报成功"})
		if err != nil {
			log.Println(err)
		}
		go s.Write(byteData)
	}
}
