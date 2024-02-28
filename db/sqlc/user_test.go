package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"gomarketplace/common"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	createUserArg := CreateUserParams{
		FirstName: common.GetRandomName(),
		LastName:  common.GetRandomName(),
		SurName:   common.GetRandomName(),
		Email:     "test1@mail.ru",
		Password:  "password",
		Address:   common.GetRandomArticleName(),
		CreatedAt: time.Now(),
	}

	user, err := testQueries.CreateUser(context.Background(), createUserArg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, createUserArg.FirstName, user.FirstName)
}
