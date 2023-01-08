package repository

import (
	"context"
	"fmt"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/models"
)

type Auth interface {
	GetByEmail(ctx context.Context, email string) error
	Create(ctx context.Context, user models.User) error
}

type repository struct {
	h db.Handler
}

func NewAuthRepository(h db.Handler) Auth {
	return &repository{h: h}
}

func (r *repository) GetByEmail(ctx context.Context, email string) error {
	query := `SELECT id FROM users WHERE email = $1`

	st, err := r.h.DB.Prepare(ctx, "getByEmail", query)
	if err != nil {
		return err
	}
	fmt.Println(st)

	row := r.h.DB.QueryRow(ctx, st.SQL, email)

	var id uint64
	return row.Scan(&id)
}

func (r *repository) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (email,password) VALUES ($1,$2)`

	st, err := r.h.DB.Prepare(ctx, "creating user", query)
	if err != nil {
		return err
	}

	_, err = r.h.DB.Exec(ctx, st.SQL, user.Email, user.Password)
	return err
}
