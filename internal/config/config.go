package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Telegram struct {
		Token string `yaml:"token"`
	} `yaml:"telegram"`
}

func LoadConfig(path string) (*Config, error) {
	// Открываем файл конфигурации
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Создаём переменную для хранения конфигурации
	var cfg Config

	// Декодируем содержимое файла YAML в структуру cfg
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
