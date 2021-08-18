package ds

type Projects struct {
	Title    string   `bson:"title"`
	Channels Channels `bson:"channels"`
}

type Channels struct {
	Vk   []VkConfig   `bson:"vk"`
	Tg   []TgConfig   `bson:"tg"`
	Jivo []JivoConfig `bson:"jivo"`
}

type VkConfig struct {
	TokenVk string `bson:"token_vk"` // Токен для ВК
}

type TgConfig struct {
	TokenTg string `bson:"token_tg"` // Токен для ТГ
}

type JivoConfig struct {
	ApiUrl      string `bson:"api_url"`      // Адрес АПИ
	WebhookPath string `bson:"webhook_path"` // Адрес нашего вебхука
}
