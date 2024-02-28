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
		fmt.Print(err)

		log.Fatal("User doesn't create")
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
	results := make(chan Product)

	for i := 0; i < n; i++ {
		go func() {
			args := CreateProductParams{
				Name: sql.NullString{String: common.GetRandomName(), Valid: true},
				Description: sql.NullString{
					String: common.GetTextMock(),
					Valid:  true,
				},
				Price: sql.NullString{
					String: "20.00",
					Valid:  true,
				},
				StockQuantity: sql.NullInt32{
					Int32: 2,
					Valid: true,
				},
				SellerID: sql.NullInt32{
					Int32: user.UserID,
					Valid: true,
				},
				CategoryID: sql.NullInt32{
					Int32: category.CategoryID,
					Valid: true,
				},
				IsActive: sql.NullBool{
					Bool:  true,
					Valid: true,
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
			result, err := store.CreateProductTx(context.Background(), args)

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
