package repositories_test

import (
	"testing"

	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
	"github.com/GoGogogogon/api/repositories/testdata"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArtileDetail(t *testing.T) {

	test := []struct {
		testTitle string
		expected  models.Article //テストの期待値
	}{
		{
			// ID1番
			testTitle: "subject",
			expected:  testdata.ArticTestData[1],
		}, {

			//ID2番
			testTitle: "subject2",
			expected:  testdata.ArticTestData[2],
		},
	}
	for _, test := range test {
		// _ には本来ループした回数が入る
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("get : %d but want %d", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("get : %s but want %s", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("get : %s but want %s", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("get : %s but want %s", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("get : %d but want %d", got.NiceNum, test.expected.NiceNum)
			}
		})
	}

}

func TestSelectArticleList(t *testing.T) {

	expectedNum := len(testdata.ArticTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	//スライスの長さが期待通りでないとき

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}

}

func TestInsertArticle(t *testing.T) {
	articles := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, articles)

	if err != nil {
		t.Error(err)
	}

	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but get %d", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlstr = `delete from articles
						where title = ? and contents = ? and username = ?
						`
		testDB.Exec(sqlstr, articles.Title, articles.Contents, articles.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {

	articlesID := 1

	// 	1. テスト結果にて期待する値を定義
	// 2. テスト対象となる関数を実行
	// 3. 2の結果と 1の値を比較
	// 一致したらテスト成功、不一致ならテスト失敗

	before, err := repositories.SelectArticleDetail(testDB, articlesID)

	if err != nil {
		t.Fatal("fail to get before")
	}

	// before 4

	// t.Cleanup(func() {
	// 	const sqrstr = `update articles
	// 					set nice = nice - 1
	// 					where article_id = ?;
	// 					`
	// 	testDB.Exec(sqrstr, articles.ID)
	// })

	err = repositories.UpdateNiceNum(testDB, articlesID)

	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articlesID)

	if err != nil {
		t.Fatal("fail to get after")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Errorf("fail to update nicenum")
	}

}
