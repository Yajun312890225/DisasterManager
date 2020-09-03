package dao

import "DisasterManager/model"

type CommunityDao model.Community

func Community() *CommunityDao {
	return &CommunityDao{}
}

func (c *CommunityDao) GetAllCommunity() (communities []CommunityDao, err error) {
	var doc []CommunityDao
	table := model.DB.Select("community.*").Table("community")
	if err := table.Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}
