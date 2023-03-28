package models

import "gorm.io/gorm"

type CreatorBase struct {
	gorm.Model
	CreatorId uint `json:"creatorId" gorm:"column:creator_id;comment:创建人 ID"`
}

type ListPageForm struct {
	Size int `comment:"每页记录数"`
	Page int `comment:"页数"`
}
