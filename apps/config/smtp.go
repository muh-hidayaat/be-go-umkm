package config

import "github.com/spf13/viper"

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func GetSMTPConfig() SMTPConfig {
	return SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: viper.GetString("SMTP_USERNAME"),
		Password: viper.GetString("SMTP_PASSWORD"),
		From:     viper.GetString("SMTP_FROM"),
	}
}
