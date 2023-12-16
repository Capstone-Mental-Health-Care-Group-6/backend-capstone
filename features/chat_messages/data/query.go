package data

import (
	root "FinalProject/features/chat_messages"
	"FinalProject/helper"
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageData struct {
	db *gorm.DB
}

func New(db *gorm.DB) root.MessageDataInterface {
	return &MessageData{
		db: db,
	}
}

func (r *MessageData) Get(chat int, query url.Values) (model []*root.Message) {
	table := r.db.Table("chat_messages").
		Where(func() string {
			format := "%s = %d AND deleted_at IS NULL"
			return fmt.Sprintf(format, []any{
				"chat_id", chat,
			}...)
		}())
	helper.QuerySorting(table, query)
	helper.QueryPagination(table, query)
	helper.QueryFiltering(table, query)
	if err := table.Find(&model).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return nil
	}
	return
}

func (r *MessageData) Find(chat, message int) (model *root.Message) {
	table := r.db.Table("chat_messages").
		Where(func() string {
			format := "%s = %d AND deleted_at IS NULL"
			return fmt.Sprintf(format, []any{
				"chat_id", chat,
			}...)
		}())
	if err := table.First(&model, message).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return nil
	}
	return
}

func (r *MessageData) Create(data *root.Message) *root.Message {
	table := r.db.Table("chat_messages")
	if err := table.Create(&data).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return nil
	}
	return r.Find(data.ChatID, data.ID)
}

func (r *MessageData) Update(data *root.Message) *root.Message {
	table := r.db.Table("chat_messages")
	search := &root.Message{ChatID: data.ChatID, ID: data.ID}
	if err := table.Find(&search).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *root.Message) *root.Message {
		fields := []string{"Text", "Blob"}
		n := structs.Map(new)
		o := structs.Map(old)
		result := make(map[string]interface{})
		for _, field := range fields {
			if n[field] != "" {
				result[field] = n[field]
			} else {
				result[field] = o[field]
			}
		}
		old.Text = result["Text"].(string)
		old.Blob = result["Blob"].(string)
		return old
	}(search, data)
	if err := table.Save(&model).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return nil
	}
	return model
}

func (r *MessageData) Delete(chat int, message int) bool {
	table := r.db.Table("chat_messages").
		Where(func() string {
			format := "%s = %d"
			return fmt.Sprintf(format, []any{
				"chat_id", chat,
			}...)
		}())
	if err := table.Delete(&root.Message{}, message).Error; err != nil {
		logrus.Error("[message.repository]: ", err.Error())
		return false
	}
	return true
}
