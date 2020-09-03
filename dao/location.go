package dao

import (
	"DisasterManager/model"
)

type LocationDao model.Location

func Location() *LocationDao {
	return &LocationDao{}
}

func (l *LocationDao) InsertLocation() error {
	if err := model.DB.Table("location").Create(l).Error; err != nil {
		return err
	}
	return nil
}

func (l *LocationDao) GetLocationByDisasterId(disasterId int, userId int64) (location []LocationDao, err error) {
	table := model.DB.Table("location").Where("disaster_id = ?", disasterId)
	if userId != 0 {
		table = table.Where("user_id = ?", userId)
	}
	if err = table.Order("created_at").Find(&location).Error; err != nil {
		return nil, err
	}
	return
}
