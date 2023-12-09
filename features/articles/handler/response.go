package handler

type InputResponse struct {
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Thumbnail  string `json:"thumbnail"`
	Status     string `json:"status"`
	Slug       string `json:"slug"`
}

type ArticleResponse struct {
	Category  string `json:"category"`
	User      string `json:"user"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	Status    string `json:"status"`
	Slug      string `json:"slug"`
}

type DashboardResponse struct {
	TotalArticle        int `json:"total_article"`
	TotalArticleBaru    int `json:"total_article_baru"`
	TotalArticlePending int `json:"total_article_pending"`
}
