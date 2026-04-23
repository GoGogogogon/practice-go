package services

import (
	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//articleのコメントに追加

	article.CommentList = append(article.CommentList, commentList...)

	return article, err
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {

	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	// if created_time.Valid {
	// 	ArticleList.CreatedAt = created_time.Time
	// }　ここいらないかも
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	ArticleList, err := repositories.SelectArticleList(s.db, page)

	if err != nil {
		return []models.Article{}, err
	}

	return ArticleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)

	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Contents:    article.Contents,
		UserName:    article.UserName,
		NiceNum:     article.NiceNum + 1,
		CommentList: article.CommentList,
		CreatedAt:   article.CreatedAt,
	}, nil
}
