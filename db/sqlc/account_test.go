package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ewallet/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		FirstName:   util.RandomFirstName(),
		LastName:    util.RandomLastName(),
		Email:       util.RandomEmail(),
		Currency:    util.RandomCurrency(),
		Balance:     util.RandomMoney(""),
		PhoneNumber: util.RandomPhoneNumber(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, arg.FirstName, account.FirstName)
	require.Equal(t, arg.LastName, account.LastName)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.PhoneNumber, account.PhoneNumber)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestReadAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.PhoneNumber, account2.PhoneNumber)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(account1.Currency),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.PhoneNumber, account2.PhoneNumber)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
