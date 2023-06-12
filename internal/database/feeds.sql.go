// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createdFeed = `-- name: CreatedFeed :one
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, created_at, updated_at, name, url, user_id
`

type CreatedFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreatedFeed(ctx context.Context, arg CreatedFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createdFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}
