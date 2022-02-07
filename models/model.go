package models

type ArticleModel struct {
	articles []Article
}

func (m *ArticleModel) AllArticles() *[]Article {
	return &m.articles
}

func (m *ArticleModel) AddArticle(a ...Article) {
	m.articles = append(m.articles, a...)
}

func NewModel() *ArticleModel {
	var articles []Article = make([]Article, 2)
	model := new(ArticleModel)
	model.articles = articles

	model.AddArticle(
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	)
	return model
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var TheModel = NewModel()
