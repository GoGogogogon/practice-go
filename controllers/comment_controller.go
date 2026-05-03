package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GoGogogogon/api/controllers/services"
	"github.com/GoGogogogon/api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostingCommentHandler(w http.ResponseWriter, req *http.Request) {

	// io.WriteString(w, "posting comment\n")

	// comment := models.Comment1
	// jsonData, err := json.Marshal(comment)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)

	//defer req.Body.Close()

	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)

	if err != nil {
		http.Error(w, "fail internal exec", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
