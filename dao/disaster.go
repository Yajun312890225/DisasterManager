package dao

import (
	"DisasterManager/model"
	"log"
)

// DisasterDao 对Disaster模型进行增删查改
type DisasterDao model.Disaster

// Disaster
func Disaster() *DisasterDao {
	return &DisasterDao{}
}

func (d *DisasterDao) GetDisasterByUser(userId int64) (disasterList []DisasterDao, err error) {
	var user model.User
	if err = model.DB.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	table := model.DB.Table("disaster").Select("disaster.*").Joins("left join disaster_community on disaster.disaster_id = disaster_community.disaster_id")
	table = table.Where("disaster_community.community_id = ? AND disaster.status = 1 ", user.CommunityId)
	if err = table.Find(&disasterList).Error; err != nil {
		return nil, err
	}
	return
}
func (d *DisasterDao) GetDisasterByCommunityId(communityId int) (disasterList []DisasterDao, err error) {
	table := model.DB.Table("disaster").Select("disaster.*").Joins("left join disaster_community on disaster.disaster_id = disaster_community.disaster_id")
	table = table.Where("disaster_community.community_id = ?", communityId)
	if err = table.Find(&disasterList).Error; err != nil {
		return nil, err
	}
	return
}

type DisasterTypeDao model.DisasterType

func DisasterType() *DisasterTypeDao {
	return &DisasterTypeDao{}
}
func (d *DisasterTypeDao) GetDisasterTypeList() (disasterTypeList []DisasterTypeDao, err error) {
	if err = model.DB.Table("disaster_type").Find(&disasterTypeList).Error; err != nil {
		return nil, err
	}
	return
}

type KnowledgeDao model.Knowledge

func Knowledge() *KnowledgeDao {
	return &KnowledgeDao{}
}

func (k *KnowledgeDao) GetKnowledgeByDisasterType(id int) (knowledgeList []KnowledgeDao, err error) {
	if err = model.DB.Table("knowledge").Where("disaster_type_id = ?", id).Find(&knowledgeList).Error; err != nil {
		return nil, err
	}
	return
}
