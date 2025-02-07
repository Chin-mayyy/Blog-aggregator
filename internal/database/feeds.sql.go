// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds(name, url, user_id)
VALUES(
    $1,
    $2,
    $3
)
RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	Name   string
	Url    string
	UserID uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed, arg.Name, arg.Url, arg.UserID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeed = `-- name: GetFeed :many
SELECT feeds.name, feeds.url, users.name FROM feeds
JOIN users ON feeds.user_id = users.id
`

type GetFeedRow struct {
	Name   string
	Url    string
	Name_2 string
}

func (q *Queries) GetFeed(ctx context.Context) ([]GetFeedRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedRow
	for rows.Next() {
		var i GetFeedRow
		if err := rows.Scan(&i.Name, &i.Url, &i.Name_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
