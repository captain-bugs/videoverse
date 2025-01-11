package repository

import (
	"videoverse/internal"
	"videoverse/pkg/logbox"
	"videoverse/pkg/models"
	"videoverse/storage"
)

type IRepo interface {
	Video() models.IVideoRepo
	User() models.IUserRepo
	Share() models.IShareRepo
	Storage() storage.IFileStore
}

type Repository struct {
	video   models.IVideoRepo
	user    models.IUserRepo
	share   models.IShareRepo
	storage storage.IFileStore
}

func NewRepository() IRepo {
	conn := internal.MakeSQLiteConnection()
	logbox.NewLogBox().Debug().Msg("setting up repository")
	return &Repository{
		video:   NewVideoRepository(conn),
		user:    NewUserRepository(conn),
		share:   NewShareRepository(conn),
		storage: storage.NewDisk(),
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

func (r *Repository) Storage() storage.IFileStore {
	return r.storage
}
