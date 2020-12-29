package model

//文章结构体--article.go
type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

//返回表名
func (a Article) TableName() string{
	return "blog_article"
}