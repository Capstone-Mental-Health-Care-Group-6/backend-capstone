package data

import (
	root "FinalProject/features/chats"
	"FinalProject/helper"
	"fmt"
	"net/url"
	"time"

	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChatData struct {
	db *gorm.DB
}

func NewChatData(db *gorm.DB) root.ChatDataInterface {
	return &ChatData{
		db: db,
	}
}

func (r *ChatData) Get(user int, query url.Values) (model []*root.Chat) {
	table := r.db.Table("chats").
		Preload("Patient").
		Preload("Doctor").
		Where(func() string {
			format := "(%s = %d OR %s = %d) AND deleted_at IS NULL"
			return fmt.Sprintf(format, []any{
				"patient_user_id", user,
				"doctor_user_id", user,
			}...)
		}())
	helper.QueryPagination(table, query)
	helper.QuerySorting(table, query)
	helper.QueryFiltering(table, query)
	if err := table.Find(&model).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return nil
	}
	return
}

func (r *ChatData) Find(chat int) (model *root.Chat) {
	table := r.db.Table("chats").
		Preload("Patient").
		Preload("Doctor").
		Where("deleted_at IS NULL")
	if err := table.First(&model, chat).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return nil
	}
	return
}

func (r *ChatData) Create(data *root.Chat) *root.Chat {
	table := r.db.Table("chats")
	if err := table.Create(&data).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return nil
	}
	return r.Find(data.ID)
}

func (r *ChatData) Update(data *root.Chat) *root.Chat {
	table := r.db.Table("chats").
		Preload("Patient").
		Preload("Doctor")
	search := &root.Chat{ID: data.ID}
	if err := table.Find(&search).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *root.Chat) *root.Chat {
		fields := []string{"LastMessage", "LastMessageTime", "LastMessageSentByID", "LastMessageSeenByID"}
		n := structs.Map(new)
		o := structs.Map(old)
		result := make(map[string]interface{})
		for _, field := range fields {
			if n[field] != nil {
				result[field] = n[field]
			} else {
				result[field] = o[field]
			}
		}
		if SentBy := result["LastMessageSentByID"].(*int); SentBy != nil {
			old.LastMessage = result["LastMessage"].(string)
			old.LastMessageTime = result["LastMessageTime"].(*time.Time)
			old.LastMessageSentByID = SentBy
			old.LastMessageSeenByID = nil
		}
		if SeenBy := result["LastMessageSeenByID"].(*int); SeenBy != nil {
			old.LastMessageSeenByID = SeenBy
		}
		return old
	}(search, data)
	if err := table.Save(&model).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return nil
	}
	return model
}

func (r *ChatData) Delete(chat int) bool {
	table := r.db.Table("chats")
	model := &root.Chat{}
	if err := table.Delete(model, chat).Error; err != nil {
		logrus.Error("[chat.repository]: ", err.Error())
		return false
	}
	return true
}
