package db

import (
	"context"
	"fmt"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/config"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB *pgx.Conn
}

func Init(ctx context.Context, c *config.Config) (Handler, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME)
	// conn, err := sql.Open("pgx", psqlInfo)
	// if err != nil {
	// 	return Handler{}, err
	// }
	conn, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		return Handler{}, err
	}

	if err := conn.Ping(ctx); err != nil {
		return Handler{}, err
	}

	return Handler{conn}, nil
}
