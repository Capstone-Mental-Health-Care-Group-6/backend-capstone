package data

import "gorm.io/gorm"

type CounselingTopic struct {
	*gorm.Model
	Name string `gorm:"column:name;type:varchar(255)"`
}
