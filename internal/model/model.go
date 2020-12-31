package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myblog/global"
	"myblog/pkg/setting"
)

//基础结构体--model.go
type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databasestring *setting.DatabaseSettings) (*gorm.DB,error) {
	db,err := gorm.Open(databasestring.DBType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s?charset=%s&parseTime=%t&loc=Local"))
	if err != nil {
		return nil,err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databasestring.MaxIdleConns)
	db.DB().SetMaxOpenConns(databasestring.MaxOpenConns)
	return db,nil
}
