package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"gomarketplace/common"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	createUserArg := CreateUserParams{
		FirstName: sql.NullString{String: common.GetRandomName(), Valid: true},
		LastName:  sql.NullString{String: common.GetRandomName(), Valid: true},
		SurName:   sql.NullString{String: common.GetRandomName(), Valid: true},
		Email:     sql.NullString{String: "test1@mail.ru", Valid: true},
		Password:  sql.NullString{String: "password", Valid: true},
		Address:   sql.NullString{String: common.GetRandomArticleName(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: false},
	}

	user, err := testQueries.CreateUser(context.Background(), createUserArg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, createUserArg.FirstName.String, user.FirstName.String)
}
