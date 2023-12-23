package data

type WordTopic struct {
	ID   int
	Name string
}

type GermanWord struct {
	ID          int
	TopicId     int
	Article     int
	Word        string
	Translation string
}

var articles []string = []string{"der", "die", "das"}

func GetArticles() []string {
	return articles
}
