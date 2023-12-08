package data

import (
	"FinalProject/features/chatbot"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatbotData struct {
	mdb *mongo.Database
}

func New(mdb *mongo.Database) chatbot.ChatbotDataInterface {
	return &ChatbotData{
		mdb: mdb,
	}
}

func (c *ChatbotData) GetAllChatBot(user_id int) ([]chatbot.Chatbot, error) {
	var result []chatbot.Chatbot
	cursor, err := c.mdb.Collection("chat_bot").Find(context.Background(), bson.M{"userid": user_id})
	if err != nil {
		return result, err
	}

	for cursor.Next(context.Background()) {
		var data chatbot.Chatbot
		err := cursor.Decode(&data)
		if err != nil {
			return result, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (c *ChatbotData) InsertChatBot(input chatbot.Chatbot) (chatbot.Chatbot, error) {
	result, err := c.mdb.Collection("chat_bot").InsertOne(context.Background(), input)
	if err != nil {
		return chatbot.Chatbot{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return chatbot.Chatbot{}, errors.New("failed to get inserted ID")
	}

	newData := chatbot.Chatbot{
		ID:           insertedID,
		UserID:       input.UserID,
		Prompt:       input.Prompt,
		ResultPrompt: input.ResultPrompt,
		Date:         input.Date,
	}

	return newData, err
}
