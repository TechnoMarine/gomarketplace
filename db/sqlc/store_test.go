package db

import (
	"context"
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
		FirstName: common.GetRandomName(),
		LastName:  common.GetRandomName(),
		SurName:   common.GetRandomName(),
		Email:     "test1@mail.ru",
		Password:  "password",
		Address:   common.GetRandomArticleName(),
		CreatedAt: time.Now(),
	}

	user, err := testQueries.CreateUser(context.Background(), createUserArg)
	if err != nil {
		log.Fatal("User doesn't create")
	}

	createCategoryArgs := CreateCategoryParams{
		CategoryName: common.GetRandomArticleName(),
		CreatedAt:    time.Time{},
	}

	category, err := testQueries.CreateCategory(context.Background(), createCategoryArgs)
	if err != nil {
		fmt.Print("{}", err)
		log.Fatal("Category doesnt created")
	}

	errs := make(chan error)
	results := make(chan Product)

	for i := 0; i < n; i++ {
		go func() {
			args := CreateProductParams{
				Name:          common.GetRandomName(),
				Description:   common.GetTextMock(),
				Price:         "20.00",
				StockQuantity: 2,
				SellerID:      user.UserID,
				CategoryID:    category.CategoryID,
				IsActive:      true,
				CreatedAt:     time.Time{},
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
