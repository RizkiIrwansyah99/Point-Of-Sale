package repository

import (
	"context"
	"database/sql"
	"github.com/arioprima/blog_web/models/entity"
)

type UserRepository interface {
	create(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	update(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	delete(context.Context, *sql.Tx, *entity.User) error
	FindById(ctx context.Context, tx *sql.Tx, userId string) (*entity.User, error)
	FindByUserName(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.User, error)
}
