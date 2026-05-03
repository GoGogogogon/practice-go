package api

import (
	"database/sql"
	"net/http"

	"github.com/GoGogogogon/api/controllers"
	"github.com/GoGogogogon/api/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	acon := controllers.NewArticleControllers(ser)
	ccon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/hello", acon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", acon.PostArticleHandler).Methods(http.MethodPost)

	r.HandleFunc("/article/list", acon.ArticleListHandler).Methods(http.MethodGet)

	r.HandleFunc("/article/{id:[0-9]+}", acon.ArticleDetailHandler).Methods(http.MethodGet)

	r.HandleFunc("/articlenice", acon.PostingNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", ccon.PostingCommentHandler).Methods(http.MethodPost)

	return r
}
