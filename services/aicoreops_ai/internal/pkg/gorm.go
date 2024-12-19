package pkg

import (
	"aicoreops_ai/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(cfg config.MySQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database),
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, nil
}
