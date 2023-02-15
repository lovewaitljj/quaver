package mysql

import (
	"quaver/models"
	"quaver/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "root1234",
		DbName:   "quaver",
		Port:     3306,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}
func TestCreatePost(t *testing.T) {
	user := models.User{
		Name:     "blingder",
		Password: "12345678",
	}
	_, err := Register(&user)
	if err != nil {
		t.Fatalf("Register insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("Register insert record into mysql success")
}
