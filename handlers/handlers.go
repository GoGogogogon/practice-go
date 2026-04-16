package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/GoGogogogon/api/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	// length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	// if err != nil {
	// 	http.Error(w, "cannot get content length\n", http.StatusBadRequest)
	// 	return
	// }

	// reqBodybuffer := make([]byte, length)

	// if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
	// 	http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	// 	return
	// }

	defer req.Body.Close() //ボディをclose
	var reqArticle models.Article
	// if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
	// 	http.Error(w, "fail to decode json", http.StatusBadRequest)
	// 	return
	// }

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return

	}

	article := reqArticle
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(jsonData)
	//io.WriteString(w, "Posting Article…\n")
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {

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

	reqArticle := []models.Article{models.Article1, models.Article2} //Article1とArticle2のスライス

	json.NewEncoder(w).Encode(reqArticle)

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

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	// articleid, err := strconv.Atoi(mux.Vars(req)["id"])
	// if err != nil {
	// 	http.Error(w, "Invaild query paramenter", http.StatusBadRequest)
	// 	return
	// }
	// resstring := fmt.Sprintf("Article No.%d\n", articleid)
	// io.WriteString(w, resstring)

	article := models.Article1

	json.NewEncoder(w).Encode(article)
	// jsonData, err := json.Marshal(article)
	// if err != nil {

	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return

	// }

	// w.Write(jsonData)

}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {

	// article := models.Article1
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)
	defer req.Body.Close()

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
}

func PostingCommentHandler(w http.ResponseWriter, req *http.Request) {

	// io.WriteString(w, "posting comment\n")

	// comment := models.Comment1
	// jsonData, err := json.Marshal(comment)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)

	defer req.Body.Close()

	var reqArticle models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
}
