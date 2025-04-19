package domain_response

type NoteResponseDomain struct {
	Status       string
	TotalResults int
	Articles     []ArticleResponseDomain
}

type ArticleResponseDomain struct {
	Source      ArticleSourceResponseDomain
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponseDomain struct {
	ID   *string
	Name string
}
