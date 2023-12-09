package chat_messages

import (
	dataUser "FinalProject/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        int            `gorm:"type:int;primaryKey;autoIncrement"`
	ChatID    int            `gorm:"type:int"`
	UserID    int            `gorm:"type:int"`
	User      dataUser.User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text      string         `gorm:"type:text"`
	Blob      string         `gorm:"type:text"`
	Timestamp time.Time      `gorm:"type:datetime(3)"`
	CreatedAt time.Time      `gorm:"type:datetime(3);autoCreateTime"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(3);index"`
}

func (Message) TableName() string {
	return "chat_messages"
}
