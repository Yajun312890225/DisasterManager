package dao

import "DisasterManager/model"

// BillBoardDao 对BillBoard模型进行增删查改
type BillBoardDao model.BillBoard

// User
func BillBoard() *BillBoardDao {
	return &BillBoardDao{}
}

// func(b *BillBoardDao)
