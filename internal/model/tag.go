package model

import "myblog/pkg/app"

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}
type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}