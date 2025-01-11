package repository

import (
	"context"
	"database/sql"
	"time"
	videoversedb "videoverse/db/videoverse"
	"videoverse/pkg/models"
)

type UserRepository struct {
	db *videoversedb.Queries
}

func NewUserRepository(connection *sql.DB) models.IUserRepo {
	return &UserRepository{db: videoversedb.New(connection)}
}

func (u UserRepository) GetByID(ctx context.Context, ID int64) (*models.User, error) {
	user, err := u.db.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	data := &models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	if user.CreatedAt.Valid {
		data.CreatedAt = user.CreatedAt.Time
	}
	return data, nil

}

func (u UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	arg := videoversedb.SaveUserParams{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: "",
		CreatedAt:    sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	outcome, err := u.db.SaveUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	data := &models.User{
		ID:       outcome.ID,
		Username: outcome.Username,
		Email:    outcome.Email,
	}
	if outcome.CreatedAt.Valid {
		data.CreatedAt = outcome.CreatedAt.Time
	}
	return data, nil
}

func (u UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}
