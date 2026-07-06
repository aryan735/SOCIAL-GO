package store

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int64     `json:"ID"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO USERS (USERNAME, EMAIL, PASSWORD) VALUES($1, $2, $3) RETURNING id, created_at`
	err := u.db.QueryRowContext(ctx, query, user.UserName, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
