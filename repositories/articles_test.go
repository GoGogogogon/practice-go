package repositories_test

import (
	"testing"

	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArtileDetail(t *testing.T) {

	//データのオープン。そもそも読み込めなければt.fatalで強制終了
	//終わったら db.closeで閉じる

	// expected := models.Article{
	// 	ID:       1,
	// 	Title:    "firstPost",
	// 	Contents: "This is my first blog",
	// 	UserName: "saki",
	// 	NiceNum:  3,
	// }

	//↑みたいに一々定義するのは
	// テストケースは増えて面倒になる

	//↓のようにまとめて構造体にする

	test := []struct {
		testTitle string
		expected  models.Article //テストの期待値
	}{
		{
			// ID1番
			testTitle: "subject",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		}, {
			//ID2番
			testTitle: "subject2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}
	//got, err := repositories.SelectArticleDetail(db, expected.ID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 		//データベース接続に失敗したときとかに使う
	// 		//その後の処理は行われない
	// 	}

	// 	if got.ID != expected.ID {
	// 		t.Errorf("get %d but want %d", got.ID, expected.ID)
	// 		//テストは失敗するが、その後の処理が続けられる
	// 	}

	// 	if got.Title != expected.Title {
	// 		t.Errorf("get %s but want %s", got.Title, expected.Title)
	// 	}

	// 	if got.Contents != expected.Contents {
	// 		t.Errorf("get %s but want %s", got.Contents, expected.Contents)
	// 	}

	// 	if got.UserName != expected.UserName {
	// 		t.Errorf("got %s but want %s", got.UserName, expected.UserName)
	// 	}

	// 	if got.NiceNum != expected.NiceNum {
	// 		t.Errorf("got %d but want %d", got.NiceNum, expected.NiceNum)
	// 	}
	// }
	//↑　一々書くのが面倒だし、量が増えたら大変
	//↓for文でリファクタリングする

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

	//データベースの準備
	//select .. detailと同じ処理をしている　非効率

	//テスト対象の関数
	expectedNum := 2
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
