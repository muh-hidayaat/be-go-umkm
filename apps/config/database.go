package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnect() *gorm.DB {
	host := viper.GetString("MYSQL_HOST")
	port := viper.GetInt("MYSQL_PORT")
	user := viper.GetString("MYSQL_USER")
	password := viper.GetString("MYSQL_PASSWORD")
	dbname := viper.GetString("MYSQL_DB")

	// Tambahkan parameter statement_timeout ke dalam dsn
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai statement_timeout=10000", // Timeout 10 detik (10000 milidetik)
	// 	host, port, user, password, dbname)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected ... !")
	return db
}
