package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct /*コメントの構造体*/ {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id" `
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct /*記事の構造体*/ {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 2,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}
	//fmt.Printf("%+v\n", article)

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%s\n", jsonData)

}
