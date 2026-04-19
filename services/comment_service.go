package services

import (
	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {

	db, err := connectDB()

	if err != nil {
		return models.Comment{}, nil
	}

	defer db.Close()

	CommentList, err := repositories.InsertComment(db, comment)

	if err != nil {
		return models.Comment{}, err
	}

	return CommentList, nil
}
