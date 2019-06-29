package story

//Story represents story record data
type Story struct {
	//Title is a title for a story
	Title string `json:"title"`
	//Content is what the story about
	Content string `json:"content"`
	//EndEmoji stores emoji that want to
	//be displayed in the end of the content
	EndEmoji string `json:"endEmoji"`
}
