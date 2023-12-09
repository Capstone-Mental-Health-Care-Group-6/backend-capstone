package chats

import (
	message "FinalProject/features/chat_messages"
	user "FinalProject/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID                  int               `gorm:"type:int;primaryKey;autoIncrement"`
	PatientUserID       int               `gorm:"type:int"`
	Patient             user.User         `gorm:"foreignKey:PatientUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DoctorUserID        int               `gorm:"type:int"`
	Doctor              user.User         `gorm:"foreignKey:DoctorUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Messages            []message.Message `gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LastMessage         string            `gorm:"type:text"`
	LastMessageTime     *time.Time        `gorm:"type:datetime(3)"`
	LastMessageSentByID *int              `gorm:"type:int"`
	LastMessageSentBy   user.User         `gorm:"foreignKey:LastMessageSentByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LastMessageSeenByID *int              `gorm:"type:int"`
	LastMessageSeenBy   user.User         `gorm:"foreignKey:LastMessageSeenByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt           time.Time         `gorm:"type:datetime(3);autoCreateTime"`
	UpdatedAt           time.Time         `gorm:"type:datetime(3);autoUpdateTime"`
	DeletedAt           gorm.DeletedAt    `gorm:"type:datetime(3);index"`
}

func (Chat) TableName() string {
	return "chats"
}
