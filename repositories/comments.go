package repositories

// 新規投稿をデータベースに insert する関数
// -> データベースに保存したコメント内容と、発生したエラーを返り値にする
import (
	"database/sql"
	"fmt"

	"github.com/GoGogogogon/api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	
	insert into comments (article_id, message, created_at) values
	(?, ?, now());
`
	var newComment models.Comment

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)

	if err != nil {
		fmt.Print(err)
		return newComment, err
	}

	newId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return newComment, err
	}

	newComment.CommentID = int(newId)
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message

	// (問 5) 構造体 models.Comment を受け取って、それをデータベースに挿入する処理
	return newComment, nil
}

// 指定 ID の記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
	select *
	from comments
	where article_id = ?;
`
	commentArray := make([]models.Comment, 0)
	result, err := db.Query(sqlStr, articleID)

	if err != nil {
		fmt.Print(err)
		return commentArray, err
	}

	defer result.Close()

	for result.Next() {
		var comment models.Comment
		var created_time sql.NullTime

		err := result.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &created_time)

		if created_time.Valid {
			comment.CreatedAt = created_time.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			commentArray = append(commentArray, comment)
		}

	}

	// (問 6) 指定 ID の記事についたコメント一覧をデータベースから取得して、
	// それを `models.Comment`構造体のスライス `[]models.Comment`に詰めて返す処理
	//return commentArray, nil
	return commentArray, err
}
