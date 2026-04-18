package repositories_test

import (
	"testing"

	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

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

	for _, got := range test {

		t.Run(got.testTitle, func(t *testing.T) {

			comments, err := repositories.SelectCommentList(testDB, got.Articleid)
			if err != nil {
				t.Fatal(err)
			}
			if num := len(comments); num != len(got.expected) {
				t.Fatalf("want %v but %v", len(got.expected), len(comments))
			}

			//test[0] = testTitle: "comment1", Articleid: 1, expected[0] models.comments { ........}
			//記事一つの中に1つ以上のコメントがあることもある。
			//全部探索するには2重for.
			for j, Getcomments := range comments {

				if Getcomments.CommentID != got.expected[j].CommentID {
					t.Errorf("want %v but %v", got.expected[j].CommentID, Getcomments.CommentID)
				}
			}

		})
	}

}

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
		t.Errorf("new commentID id is expected %d but get %d", expectedCommentsnum, got.CommentID)
	}

	t.Cleanup(func() {
		const sqlstr = `delete  from comments
						where article_id = ? and message = ?;
						`
		testDB.Exec(sqlstr, comments.ArticleID, comments.Message)
	})

}
