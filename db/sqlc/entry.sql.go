// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: entry.sql

package db

import (
	"context"
	"database/sql"
)

const createEntry = `-- name: CreateEntry :one
insert into entries (
    account_id,
    amount
) values (
             $1, $2
         ) returning id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID sql.NullInt64 `json:"account_id"`
	Amount    int64         `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
select id, account_id, amount, created_at from entries
where id = $1 limit 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entries, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
select id, account_id, amount, created_at from entries
where account_id = $1
order by id
    limit $2 offset  $3
`

type ListEntriesParams struct {
	AccountID sql.NullInt64 `json:"account_id"`
	Limit     int32         `json:"limit"`
	Offset    int32         `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entries
	for rows.Next() {
		var i Entries
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
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
