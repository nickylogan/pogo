package configs

import (
	"os"
)

// New function returns a new config function filled with environment variables.
func New() *Config {
	return &Config{
		Chat: ChatConfig{
			MasterID:           getEnv("MASTER_ID", ""),
			ChannelAccessToken: getEnv("LINE_ACCESS_TOKEN", ""),
			ChannelSecret:      getEnv("LINE_CHANNEL_SECRET", ""),
		},
		Interpretor: InterpretorConfig{
			Token: getEnv("WIT_TOKEN_KEY", ""),
		},
		Port:     getEnv("PORT", "8080"),
		CertFile: getEnv("CERT_FILE", "https-server.crt"),
		KeyFile:  getEnv("KEY_FILE", "https-server.key"),
		Weather: WeatherConfig{
			APIKey: getEnv("WEATHER_API_KEY", ""),
		},
		News: NewsConfig{
			APIKey: getEnv("NEWS_API_KEY", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
