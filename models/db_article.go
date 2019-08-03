package models

import (
	"errors"
	"strconv"
)

// slice of pointer, each element is a pointer pointing to the address of an article object
// var articles []*Article

var articles []Article
var article Article

func AddArticle(article Article) error {
	for _, art := range articles {
		if article.ID == art.ID {
			return errors.New("article alreay exists")
		}
	}
	articles = append(articles, article)
	return nil
}

// func GetArticles() ([]*Article, error) {
func GetArticles() ([]Article, error) {
	if len(articles) == 0 {
		for i := 0; i < 3; i++ {
			// articles = append(articles, &Article{i + 1, "title" + strconv.Itoa(i+1), "content" + strconv.Itoa(i+1)})
			articles = append(articles, Article{i + 1, "title" + strconv.Itoa(i+1), "content" + strconv.Itoa(i+1)})
		}
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article *Article
	for i, _ := range articles {
		if articles[i].ID == id {
			article = &articles[i] //to update the element, use address and index
		}
	}
	if article == nil {
		return article, errors.New("article not found")
	}
	return article, nil
}

func UpdateArticle(id int, article Article) error {
	art, err := GetArticle(id)
	if err != nil {
		return err
	}
	// art = &article // error, art declared and not used
	// art.ID = article.ID // don't change ID
	art.Title = article.Title
	art.Content = article.Content
	return nil
}

func DeleteArticle(id int) error {
	var idx int = -1
	for i, _ := range articles {
		if articles[i].ID == id {
			idx = i
		}
	}
	if idx == -1 {
		return errors.New("article not found")
	}
	articles = append(articles[:idx], articles[idx+1:]...) //order kept
	return nil
}
