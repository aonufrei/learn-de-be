package data

type Repository struct {
	Topics []WordTopic
	Words  []GermanWord
}

func CreateRepository(excelFilepath string) *Repository {
	topics, words, err := RefreshFromExel(excelFilepath)
	if err != nil {
		panic("Failed to get words from excel file" + err.Error())
	}
	return &Repository{
		Topics: topics,
		Words:  words,
	}
}

func (r *Repository) GetAllTopics() []WordTopic {
	return r.Topics[:]
}

func (r *Repository) GetWordsByTopic(topicId int) []GermanWord {
	words := make([]GermanWord, 0)
	for _, w := range r.Words {
		if w.TopicId != topicId {
			continue
		}
		words = append(words, w)
	}
	return words
}
