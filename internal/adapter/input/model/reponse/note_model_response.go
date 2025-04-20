package model_response

type NoteResponseModel struct {
	Status       string                 `json:"status"`
	TotalResults int                    `json:"total_results"`
	Articles     []ArticleModelResponse `json:"articles"`
}

type ArticleModelResponse struct {
	Source      ArticleSourceModelResponse `json:"source"`
	Author      string                     `json:"author"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	URL         string                     `json:"url"`
	URLToImage  string                     `json:"url_to_image"`
	PublishedAt string                     `json:"published_at"`
	Content     string                     `json:"content"`
}

type ArticleSourceModelResponse struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}
