package routers

import (
	"net/http"

	"github.com/GoGogogogon/api/controllers"
	"github.com/gorilla/mux"
)

func Router(con *controllers.MyAppController) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)

	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)

	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)

	r.HandleFunc("/articlenice", con.PostingNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", con.PostingCommentHandler).Methods(http.MethodPost)

	return r
}
