package repository

import (
	"context"
	"fmt"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/models"
)

type Repository struct {
	H db.Handler
}

func NewAuthRepository(h db.Handler) *Repository {
	return &Repository{H: h}
}

func (r *Repository) GetByEmail(ctx context.Context, email string) error {
	query := `SELECT id FROM users WHERE email = $1`

	st, err := r.H.DB.Prepare(ctx, "getByEmail", query)
	if err != nil {
		return err
	}
	fmt.Println(st)

	row := r.H.DB.QueryRow(ctx, st.SQL, email)

	// row := st.QueryRowContext(ctx, email)

	var id uint64
	return row.Scan(&id)
}

func (r *Repository) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (email,password) VALUES ($1,$2)`
	fmt.Println(1)

	st, err := r.H.DB.Prepare(ctx, "creating user", query)
	// st, err := r.H.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	fmt.Println(1)
	// defer st.Close()

	_, err = r.H.DB.Exec(ctx, st.SQL, user.Email, user.Password)
	// _, err = st.ExecContext(ctx, user.Email, user.Password)
	return err
}