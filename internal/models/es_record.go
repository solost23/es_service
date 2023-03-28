package models

import "gorm.io/gorm"

type ESRecord struct {
	CreatorBase
	DocumentID string `json:"document_id" gorm:"column:document_id;type:varchar(300);comment: 文档ID"`
	Document   string `json:"document" gorm:"column:document;type:text;comment: 文档内容"`
	Index      string `json:"index" gorm:"column:index;type:varchar(300);comment: 文档索引名"`
	Type       string `json:"type" gorm:"column:type;type:varchar(300);comment: 文档类型"`
	Result     string `json:"result" gorm:"column:result;type:varchar(300);comment: 操作结果"`
	Status     int    `json:"status" gorm:"column:status;type:bigint;comment: 状态"`
}

func (t *ESRecord) TableName() string {
	return "es_records"
}

func (t *ESRecord) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}
