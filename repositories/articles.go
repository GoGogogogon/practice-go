package repositories

import (
	"database/sql"
	"fmt"

	"github.com/GoGogogogon/api/models"
)

// POST /article

// リクエストボディで受け取った記事を投稿する
// – 構造体 models.Article を受け取って、それをデータベースに挿入する処理が必要

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?,?,?,0,now());
	`

	var newArticle models.Article

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return newArticle, err
	}

	newId, err := result.LastInsertId()
	if err != nil {
		return newArticle, err
	}

	newArticle.ID = int(newId)
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName
	newArticle.NiceNum = article.NiceNum

	return newArticle, nil

	// 結果を確認
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())
}

// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {

	const limit = 5
	const sqlStr = `
	select article_id, title, contents, username, nice
	from articles
	limit ? offset ?;
	`
	//var articleArray []models.Article

	articleArray := make([]models.Article, 0)

	result, err := db.Query(sqlStr, limit, (page-1)*limit)

	if err != nil {
		fmt.Println(err)
		return articleArray, err
	}
	defer result.Close()

	for result.Next() {

		var article models.Article

		err := result.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
			return articleArray, err
		} else {
			articleArray = append(articleArray, article)
		}

	}
	return articleArray, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	var article models.Article
	const sqlStr = `
	select *
	from articles
	where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)

	if err := row.Err(); err != nil {
		fmt.Print(err)
		return article, err
	}

	var created_time sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &created_time)

	if err != nil {
		fmt.Println(err)
		return article, err
	}

	if created_time.Valid {
		article.CreatedAt = created_time.Time
	}

	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)

	if err != nil {
		fmt.Print(err)
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
