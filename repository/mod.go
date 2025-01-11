package repository

import "videoverse/pkg/models"

type IRepo interface {
	Video() models.IVideoRepo
	User() models.IUserRepo
	Share() models.IShareRepo
}

type Repository struct {
	video models.IVideoRepo
	user  models.IUserRepo
	share models.IShareRepo
}

func NewRepository() IRepo {
	return &Repository{
		video: NewVideoRepository(),
		user:  NewUserRepository(),
		share: NewShareRepository(),
	}
}

func (r *Repository) Video() models.IVideoRepo {
	return r.video
}

func (r *Repository) User() models.IUserRepo {
	return r.user
}

func (r *Repository) Share() models.IShareRepo {
	return r.share
}
