package model

import (
	"github.com/jinzhu/gorm"
	"myblog/pkg/app"
)

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}
func (t Tag) TableName() string {
	return "blog_tag"
}
type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag) Count(db *gorm.DB) (int,error){
	var count int
	if t.Name != ""{
		db = db.Where("name = ?",t.Name)
	}
	db = db.Where("state = ?",t.State)
	if err:= db.Model(&t).Where("is_del = ?",0).Count(&count).Error;err!=nil{
		return 0,err
	}
	return count,nil
}

func (t Tag) List(db *gorm.DB,pageOffset,PageSize int) ([]*Tag,error){
	var tag []*Tag
	var err error
	if pageOffset >=0 && PageSize >0 {
		db = db.Offset(pageOffset).Limit(PageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}
	db = db.Where("state = ?",t.State)
	if err = db.Where("is_del = ?",0).Find(&tag).Error;err!=nil{
		return nil,err
	}
	return tag,nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB,values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del = ?",t.ID,0).Update(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?",t.Model.ID,0).Delete(&t).Error
}