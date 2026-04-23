package services

import (
	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	CommentList, err := repositories.InsertComment(s.db, comment)

	if err != nil {
		return models.Comment{}, err
	}

	return CommentList, nil
}
