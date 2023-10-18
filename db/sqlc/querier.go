// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Accounts, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Accounts, error)
	GetEntry(ctx context.Context, id int64) (Entries, error)
	GetSession(ctx context.Context, username string) (Sessions, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (Users, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error)
	TransferList(ctx context.Context, arg TransferListParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error)
}

var _ Querier = (*Queries)(nil)
