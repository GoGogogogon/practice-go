package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Print(err)
		return
	}
	// if err := db.Ping(); err != nil {
	// 	fmt.Print(err)
	// } else {
	// 	fmt.Print("connnect to DB")
	// }/*

	//データ挿入

	// article := models.Article{
	// 	Title:    "insert test",
	// 	Contents: "can i insert data correctly",
	// 	UserName: "saki",
	// }

	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
		`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Print(err)
		tx.Rollback()
		return
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `
		update articles
		set nice = ?
		where article_id = ?
		`

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)

	if err != nil {
		fmt.Print(err)
		tx.Rollback()
		return
	}

	tx.Commit()

	// const sqlStr = `insert into articles(title, contents, username, nice,created_at) values
	// 				(?,?,?,0,now());
	// 				`

	//result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	//var createdTime sql.NullTime
	// err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CrearedAt)

	// err = result.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	// if err != nil {
	// 	fmt.Print(err)
	// }
	// if createdTime.Valid {
	// 	article.CrearedAt = createdTime.Time
	// }

}
