package pkg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库
func InitDB(addr string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	if err = InitTables(db); err != nil {
		panic(err)
	}

	return db
}

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate()
}
