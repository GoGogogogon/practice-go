package services

import (
	"encoding/json"
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
