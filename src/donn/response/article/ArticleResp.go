package article

import (
	"donn/entity"
)

type ArticleResp struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func MakeShowArticleResp(article *entity.Article) ArticleResp {
	return ArticleResp{article.Id, article.Title, article.Content}
}

func MakeIndexArticleResp(articles []entity.Article) []ArticleResp {
	var respSlice = []ArticleResp{}
	for _, element := range articles {
		respSlice = append(respSlice, ArticleResp{element.Id, element.Title, element.Content})
	}
	return respSlice
}
