package handler

type InputRequest struct {
	CategoryID uint   `json:"category_id" form:"category_id"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Thumbnail  string `json:"thumbnail" form:"thumbnail"`
}

type UpdateRequest struct {
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}
