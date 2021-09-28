package ds

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectString struct {
	ID       string   `json:"_id"`
	Title    string   `bson:"title"`
	Channels Channels `bson:"channels"`
}

type JustID struct {
	ID string `json:"_id"`
}

type Project struct {
	ID       primitive.ObjectID `bson:"_id"`
	Title    string             `bson:"title"`
	Channels Channels           `bson:"channels"`
}

type Channels struct {
	Vk   []VkConfig   `bson:"vk"`
	Tg   []TgConfig   `bson:"tg"`
	Jivo []JivoConfig `bson:"jivo"`
}

type VkConfig struct {
	Token string `bson:"token"` // Токен для ВК
}

type TgConfig struct {
	Token string `bson:"token"` // Токен для ТГ
}

type JivoConfig struct {
	ApiUrl      string `bson:"api_url"`      // Адрес АПИ
	WebhookPath string `bson:"webhook_path"` // Адрес нашего вебхука
}
