package response

type GetNewsResponse []GetNewsResponseItem

type GetNewsResponseItem struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Img       string `json:"img"`
	ShortDesc string `json:"short_desc"`
}
