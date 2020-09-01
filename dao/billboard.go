package dao

import "DisasterManager/model"

// BillBoardDao 对BillBoard模型进行增删查改
type BillBoardDao model.BillBoard

// BillBoard
func BillBoard() *BillBoardDao {
	return &BillBoardDao{}
}

func (b *BillBoardDao) GetBillBoard() ([]BillBoardDao, error) {
	var doc []BillBoardDao
	table := model.DB.Select("bill_board.*").Table("bill_board").Where("status = 1")
	if err := table.Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}
