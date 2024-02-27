package db

import (
	"context"
	"database/sql"
	"fmt"
	"gomarketplace/common"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateProductTx(t *testing.T) {
	store := NewStore(testConn)

	n := 10

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
	if err != nil {
		log.Fatal("User doesnt create")
	}

	createCategotyArg := CreateCategoryParams{
		CategoryName: sql.NullString{String: common.GetRandomArticleName(), Valid: true},
		ParentID: sql.NullInt32{
			Int32: 0,
			Valid: false,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}

	category, err := testQueries.CreateCategory(context.Background(), createCategotyArg)
	if err != nil {
		log.Fatal("Category doesnt created")
	}

	errs := make(chan error)
	results := make(chan Post)

	for i := 0; i < n; i++ {
		go func() {
			args := CreatePostTxParams{
				Title:  common.GetRandomArticleName(),
				Body:   common.GetTextMock(),
				UserId: user.ID,
				Status: common.GetRandomStatuses(),
			}
			result, err := store.CreatePostTx(context.Background(), args)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)

		fmt.Print(result)
	}

}
