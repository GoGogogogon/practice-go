package services

import "github.com/GoGogogogon/api/models"

//articleのみ
type ArticleServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

//commentのみ
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
