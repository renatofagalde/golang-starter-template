package response

type NoteResponseModel struct {
	Status       string
	TotalResults int
	Articles     []ArticleResponseModel
}

type ArticleResponseModel struct {
	Source      ArticleSourceModel
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceModel struct {
	ID   *string
	Name string
}
