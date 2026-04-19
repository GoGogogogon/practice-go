package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
	"github.com/gorilla/mux"
)

func ArticleDatailHandler(w http.ResponseWriter, req *http.Request) {

	articleID := mux.Vars(req)["id"]

	log.Println(articleID)

	article := models.Article1

	json.NewEncoder(w).Encode(article)
}

func GetArticService(articleID int) (models.Article, error) {

	db, err := conenctDB()
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

	var created_time sql.NullTime

	db, err := conenctDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	ArticleList, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	if created_time.Valid {
		ArticleList.CreatedAt = created_time.Time
	}

	return ArticleList, nil
}

func GetArticListHandler(page int) ([]models.Article, error) {
	db, err := conenctDB()
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

func PostNiceHandler(article models.Article) (models.Article, error) {

	var ArticleList models.Article

	db, err := conenctDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	before, err := repositories.SelectArticleDetail(db, article.ID)
	if err != nil {
		fmt.Println("fail to get before")
	}

	err = repositories.UpdateNiceNum(db, article.ID)

	if err != nil {
		return models.Article{}, err
	}

	after, err := repositories.SelectArticleDetail(db, article.ID)

	if err != nil {
		fmt.Println("fail to get after")
		return models.Article{}, err
	}

	if after.NiceNum-before.NiceNum != 1 {
		fmt.Print("fail to up nicenum")
		return models.Article{}, err
	}

	ArticleList.NiceNum = after.NiceNum

	return ArticleList, nil
}
