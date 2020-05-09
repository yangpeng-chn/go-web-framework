package models

import (
	"errors"
	"fmt"
	"strconv"
)

// slice of pointer, each element is a pointer pointing to the address of an article object
// var articles []*Article

var articles []Article
var article Article

// AddArticle adds an article into the slice
func (a *Article) AddArticle() error {
	for _, art := range articles {
		if a.ID == art.ID {
			return errors.New("article alreay exists")
		}
	}
	articles = append(articles, *a)
	return nil
}

// GetArticles returns all the articles
func (a *Article) GetArticles() ([]Article, error) {
	if len(articles) == 0 {
		for i := 0; i < 3; i++ {
			articles = append(articles, Article{i + 1, "title" + strconv.Itoa(i+1), "content" + strconv.Itoa(i+1)})
		}
	}
	return articles, nil
}

// GetArticle return an article by ID
func (a *Article) GetArticle(id int) (*Article, error) {
	for i := range articles {
		fmt.Println(id)
		if articles[i].ID == id {
			a = &articles[i] //to update the element, use address and index
		}
	}
	if a == nil {
		return &Article{}, errors.New("article not found")
	}
	return a, nil
}

// UpdateArticle updates an article by ID
func (a *Article) UpdateArticle(id int) error {
	art, err := a.GetArticle(id)
	if err != nil {
		return err
	}
	// art = &article // error, art declared and not used
	// art.ID = article.ID // don't change ID
	art.Title = a.Title
	art.Content = a.Content
	return nil
}

// DeleteArticle delete an article by ID
func (a *Article) DeleteArticle(id int) error {
	var idx int = -1
	for i := range articles {
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
