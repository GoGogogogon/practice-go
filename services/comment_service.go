package services

import (
	"github.com/GoGogogogon/api/models"
	"github.com/GoGogogogon/api/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {

	var Comment models.Comment
	db, err := conenctDB()

	if err != nil {
		return models.Comment{}, nil
	}

	defer db.Close()

	CoomentList, err := repositories.InsertComment(db, Comment)

	if err != nil {
		return models.Comment{}, err
	}

	return CoomentList, nil
}
