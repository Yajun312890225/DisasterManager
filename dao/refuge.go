package dao

import (
	"DisasterManager/model"
	"log"
)

type RefugeDao model.Refuge

func Refuge() *RefugeDao {
	return &RefugeDao{}
}

type RefugeFacilityDao model.RefugeFacility

func RefugeFacility() *RefugeFacilityDao {
	return &RefugeFacilityDao{}
}

func (r *RefugeDao) GetRefugeByUserId(userId int64) (refugeList []RefugeDao, err error) {
	var user model.User
	if err = model.DB.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	table := model.DB.Table("refuge").Select("refuge.*")
	table = table.Where("community_id = ? ", user.CommunityId)
	if err = table.Find(&refugeList).Error; err != nil {
		return nil, err
	}
	return
}
func (r *RefugeFacilityDao) GetRefugeFacilityByUserId(userId int64) (refugeFacilityList []RefugeFacilityDao, err error) {
	var user model.User
	if err = model.DB.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	table := model.DB.Table("refuge_facility").Select("refuge_facility.*")
	table = table.Where("community_id = ? ", user.CommunityId)
	if err = table.Find(&refugeFacilityList).Error; err != nil {
		return nil, err
	}
	return
}
