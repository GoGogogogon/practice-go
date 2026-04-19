package services

import (
	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {

	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//articleのコメントに追加

	article.CommentList = append(article.CommentList, commentList...)

	return article, err
}

func PostArticleService(article models.Article) (models.Article, error) {

	//var created_time sql.NullTime

	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	ArticleList, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	// if created_time.Valid {
	// 	ArticleList.CreatedAt = created_time.Time
	// }　ここいらないかも
	return ArticleList, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()

	ArticleList, err := repositories.SelectArticleList(db, page)

	if err != nil {
		return []models.Article{}, err
	}

	return ArticleList, nil
}

func PostNiceService(article models.Article) (models.Article, error) {

	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)

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
