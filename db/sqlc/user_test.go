package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/zaenalarifin12/simplebank/utils"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) Users {

	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.True(t, user2.PasswordChangeAt.IsZero())
	require.NotZero(t, user2.CreatedAt)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

//func TestUpdateAccount(t *testing.T) {
//	account1 := createRandomAccount(t)
//	arg := UpdateAccountParams{ID: account1.ID, Balance: utils.RandomMoney()}
//
//	account2, err := testQueries.UpdateAccount(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//	require.Equal(t, account1.ID, account2.ID)
//	require.Equal(t, account1.Owner, account2.Owner)
//	require.Equal(t, arg.Balance, account2.Balance)
//	require.Equal(t, account1.Currency, account2.Currency)
//
//	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
//}
