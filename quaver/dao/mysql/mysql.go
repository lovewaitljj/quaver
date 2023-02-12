package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"quaver/settings"
)

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		IgnoreRecordNotFoundError: true, // 忽略ErrRecordNotFound（记录未找到）错误
	})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	//db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	return
}

func Close() {
	d, err := db.DB()
	if err != nil {
		zap.L().Error("close DB failed", zap.Error(err))
		return
	}
	d.Close()
}
