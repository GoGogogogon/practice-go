package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/GoGogogogon/api/controllers/services"
	"github.com/GoGogogogon/api/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleControllers(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return

	}
	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {

	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	log.Print(page)

	//reqArticle := []models.Article{models.Article1, models.Article2} //Article1とArticle2のスライス

	articleList, err := c.service.GetArticleListService(page)

	if err != nil {
		http.Error(w, "fail internal exec", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)

	// resString := fmt.Sprintf("Article List (page %d)\n", page)
	// io.WriteString(w, resString)

	// article := []models.Article{models.Article1, models.Article2}
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)

}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	articleid, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invaild query paramenter", http.StatusBadRequest)
		return
	}
	// resstring := fmt.Sprintf("Article No.%d\n", articleid)
	// io.WriteString(w, resstring)

	// log.Print(articleid)
	// article := models.Article1

	// articleに指定したIDのarticleのないっ用を代入する

	article, err := c.service.GetArticleService(articleid)

	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
	// jsonData, err := json.Marshal(article)
	// if err != nil {

	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return

	// }

	// w.Write(jsonData)

}

func (c *ArticleController) PostingNiceHandler(w http.ResponseWriter, req *http.Request) {

	// article := models.Article1
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)
	// defer req.Body.Close()

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(article)
}
