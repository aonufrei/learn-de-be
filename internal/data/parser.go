package data

import (
	"log"

	"github.com/thedatashed/xlsxreader"
)

type wordExelRow struct {
	Topic       int
	Article     string
	Word        string
	Translation string
}

func RefreshFromExel(excel string) ([]WordTopic, []GermanWord, error) {
	log.Println("Opening exel file with words")
	xl, err := xlsxreader.OpenFile(excel)
	defer xl.Close()
	if err != nil {
		log.Println("Failed to open file")
		return []WordTopic{}, []GermanWord{}, err
	}
	topics := make([]WordTopic, 0)
	words := make([]GermanWord, 0)

	log.Println("Processing data in excel")
	topicsMap := make(map[string]bool, 0)
	for row := range xl.ReadRows(xl.Sheets[0]) {
		if len(row.Cells) != 4 {
			continue
		}
		// detect topic
		topic := row.Cells[0].Value
		if _, ok := topicsMap[topic]; !ok {
			nCategory := WordTopic{ID: len(topics), Name: topic}
			topics = append(topics, nCategory)
			topicsMap[topic] = true
		}
		// detect article
		article, err := GetArticleByName(row.Cells[1].Value)
		if err != nil {
			continue
		}
		nWord := GermanWord{
			ID:          len(words),
			TopicId:     len(topics) - 1,
			Article:     article,
			Word:        row.Cells[2].Value,
			Translation: row.Cells[3].Value,
		}

		words = append(words, nWord)
	}
	return topics, words, nil
}
