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
		Articleid int
		testTitle string
		expected  []models.Comment
	}{
		{

			testTitle: "comment1",
			Articleid: 1,
			expected: []models.Comment{
				{CommentID: 1,

					Message: "1st comment yeah"},
			},
		}, {
			testTitle: "comment2",
			Articleid: 2,
			expected: []models.Comment{
				{CommentID: 2,
					Message: "welcome"},
			},
		},
	}

	for i := 0; i < len(test); i++ {
		t.Run(test[i].testTitle, func(t *testing.T) {
			comments, err := repositories.SelectCommentList(testDB, test[i].Articleid)
			if err != nil {
				t.Fatal(err)
			}
			if num := len(comments); num != len(test[i].expected) {
				t.Fatal(err)
			}

			if comments[i].CommentID != test[i].expected[i].CommentID {
			}
		})
	}

}
