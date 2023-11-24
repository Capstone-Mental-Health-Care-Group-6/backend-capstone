package handler

type InputResponse struct {
	CategoryID uint
	AdminID    uint
	Title      string
	Content    string
	Thumbnail  string
	Status     string
	Slug       string
}

type ArticleResponse struct {
	Category  string
	User      string
	Title     string
	Content   string
	Thumbnail string
	Status    string
	Slug      string
}
