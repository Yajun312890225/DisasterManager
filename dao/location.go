package dao

import "DisasterManager/model"

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
