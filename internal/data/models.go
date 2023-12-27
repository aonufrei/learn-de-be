package data

type WordTopic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GermanWord struct {
	ID          int    `json:"id"`
	TopicId     int    `json:"topic_id"`
	Article     int    `json:"article"`
	Word        string `json:"word"`
	Translation string `json:"translation"`
}

var articles []string = []string{"der", "die", "das"}

func GetArticles() []string {
	return articles
}
