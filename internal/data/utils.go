package data

import (
	"errors"
	"strconv"
)

func GetArticleByName(name string) (int, error) {
	for i, v := range articles {
		if v == name {
			return i, nil
		}
	}
	return -1, errors.New("No article found with name [" + name + "]")
}

func GetArticleById(id int) (string, error) {
	if id < 0 || id >= len(articles) {
		return "", errors.New("No article found with id [" + strconv.Itoa(id) + "]")
	}
	return articles[id], nil
}
