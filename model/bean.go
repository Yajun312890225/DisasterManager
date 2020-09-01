package model

import "time"

// Model 基础类型
type Model struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// User 用户
type User struct {
	UserId      int64  `gorm:"primary_key;AUTO_INCREMENT:false"  json:"userId"`
	Username    string `gorm:"type:varchar(64)" json:"username" `
	Phone       string `gorm:"type:varchar(11)" json:"phone" `
	Password    string `gorm:"type:varchar(128)" json:"-"`
	Avatar      string `gorm:"type:varchar(255)" json:"avatar"`
	Sex         string `gorm:"type:varchar(255)" json:"sex"`
	Email       string `gorm:"type:varchar(128)" json:"email" `
	Birthday    string `gorm:"type:varchar(32)" json:"birthday"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	CommunityId int    `gorm:"type:int(11);" json:"communityId" `
	Remark      string `gorm:"type:varchar(255)" json:"remark"`
	Model
}

// BillBoard 宣传栏
type BillBoard struct {
	BillId int    `gorm:"primary_key;AUTO_INCREMENT"  json:"billId"`
	Title  string `gorm:"type:varchar(255)" json:"title"`
	Image  string `gorm:"type:varchar(255)" json:"image"`
	Url    string `gorm:"type:varchar(255)" json:"url"`
	Status string `gorm:"type:int(1);DEFAULT:1;" json:"status"`
	Model
}

// DisasterType 灾害类型
type DisasterType struct {
	DisasterTypeId int    `gorm:"primary_key;AUTO_INCREMENT"  json:"disasterTypeId"`
	Description    string `gorm:"type:varchar(255)" json:"description"`
	Model
}

// Knowledge 宣传知识
type Knowledge struct {
	KnowledgeId    int `gorm:"primary_key;AUTO_INCREMENT"  json:"knowledgeId"`
	DisasterTypeId int `gorm:"type:int(11);" json:"disasterTypeId" `
	Model
}

// Community 社区
type Community struct {
	CommunityId int    `gorm:"primary_key;AUTO_INCREMENT"  json:"communityId"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	Model
}

// Refuge 避难所
type Refuge struct {
	RefugeId    int     `gorm:"primary_key;AUTO_INCREMENT"  json:"refugeId"`
	CommunityId int     `gorm:"primary_key;AUTO_INCREMENT:false"  json:"communityId"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Model
}

// RefugeFacility 避难设施
type RefugeFacility struct {
	RefugeFacilityId int     `gorm:"primary_key;AUTO_INCREMENT"  json:"refugeFacilityId"`
	CommunityId      int     `gorm:"primary_key;AUTO_INCREMENT:false"  json:"communityId"`
	Longitude        float64 `json:"longitude"`
	Latitude         float64 `json:"latitude"`
	Model
}

// Disaster 灾害
type Disaster struct {
	DisasterId     int       `gorm:"primary_key;AUTO_INCREMENT"  json:"disasterId"`
	DisasterTypeId int       `gorm:"type:int(11);" json:"disasterTypeId"`
	Status         string    `gorm:"type:int(1);DEFAULT:1;" json:"status" `
	Level          int       `json:"level"`
	BeginTime      time.Time `json:"beginTime"`
	EndTime        time.Time `json:"endTime"`
	Longitude      float64   `json:"longitude"`
	Latitude       float64   `json:"latitude"`
	Remark         string    `gorm:"type:varchar(255)" json:"remark"`
	Model
}

// DisasterCommunity 灾害影响到的社区
type DisasterCommunity struct {
	DisasterId  int `gorm:"type:int(11)" json:"disasterId"`
	CommunityId int `gorm:"type:int(11)" json:"communityId"`
}

// Location 实时位置
type Location struct {
	UserId     int64     `gorm:"primary_key;AUTO_INCREMENT:false" json:"userId"`
	DisasterId int       `json:"disasterId"`
	Longitude  float64   `json:"longitude"`
	Latitude   float64   `json:"latitude"`
	CreatedAt  time.Time `gorm:"primary_key;AUTO_INCREMENT:false" json:"createdAt"`
}
