package model_response

type NoteResponseModel struct {
	Status       string                 `json:"status,omitempty"`
	TotalResults int                    `json:"total_results,omitempty"`
	Articles     []ArticleModelResponse `json:"articles,omitempty"`
}

type ArticleModelResponse struct {
	Source      ArticleSourceModelResponse `json:"source,omitempty"`
	Author      string                     `json:"author,omitempty"`
	Title       string                     `json:"title,omitempty"`
	Description string                     `json:"description,omitempty"`
	URL         string                     `json:"url,omitempty"`
	URLToImage  string                     `json:"url_to_image,omitempty"`
	PublishedAt string                     `json:"published_at,omitempty"`
	Content     string                     `json:"content,omitempty"`
}

type ArticleSourceModelResponse struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}
