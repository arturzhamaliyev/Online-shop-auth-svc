package repository

import (
	"context"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/models"
)

type Auth interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user models.User) (int64, error)
}

type repository struct {
	h db.Handler
}

func NewAuthRepository(h db.Handler) Auth {
	return &repository{h: h}
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	st, err := r.h.DB.Prepare(ctx, "getByEmail", query)
	if err != nil {
		return nil, err
	}

	row := r.h.DB.QueryRow(ctx, st.SQL, email)

	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Create(ctx context.Context, user models.User) (int64, error) {
	query := `INSERT INTO users (email,password) VALUES ($1,$2) RETURNING id`

	st, err := r.h.DB.Prepare(ctx, "creating user", query)
	if err != nil {
		return 0, err
	}

	row := r.h.DB.QueryRow(ctx, st.SQL, user.Email, user.Password)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
