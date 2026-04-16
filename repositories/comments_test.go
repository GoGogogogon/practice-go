package repositories_test

import (
	"testing"

	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func TestInsertComment(t *testing.T) {

	comments := models.Comment{
		ArticleID: 3,
		Message:   "Hello",
	}

	expectedCommentsnum := 3
	got, err := repositories.InsertComment(testDB, comments)

	if err != nil {
		t.Fatal(err)
	}

	if got.CommentID != expectedCommentsnum {
		t.Errorf("new article id is expected %d but get %d", expectedCommentsnum, got.CommentID)
	}

	t.Cleanup(func() {
		const sqlstr = `delete  from comments
						where article_id = ? and message = ?;
						`
		testDB.Exec(sqlstr, comments.ArticleID, comments.Message)
	})

}

func TestSelectCommentList(t *testing.T) {

	test := []struct {
		testTitle string
		expected  models.Comment
	}{}
}
