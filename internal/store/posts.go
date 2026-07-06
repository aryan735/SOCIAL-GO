package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type PostStore struct {
	db *sql.DB
}

func (p *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO POSTS (CONTENT, TITLE, USERID, TAGS) VALUES($1, $2, $3, $4) RETURNING ID, CREATED_AT, UPDATED_AT`
	err := p.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserId, pq.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
