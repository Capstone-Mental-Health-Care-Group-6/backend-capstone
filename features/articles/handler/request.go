package handler

type InputRequest struct {
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Thumbnail  string `json:"thumbnail"`
	Slug       string `json:"slug"`
}

type UpdateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	Slug      string `json:"slug"`
}
